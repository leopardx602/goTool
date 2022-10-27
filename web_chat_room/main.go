package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type Connection struct {
	MsgChan  chan Message
	WebConn  *websocket.Conn
	ClientIP string
}

func NewConnection(ctx *gin.Context) (connection *Connection, err error) {
	var upgrader = websocket.Upgrader{}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed to Upgrade: ", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	return &Connection{
		MsgChan:  make(chan Message, 128),
		WebConn:  conn,
		ClientIP: ctx.ClientIP(),
	}, nil
}

func (c *Connection) GetMsg() Message {
	return <-c.MsgChan
}
func (c *Connection) AddMsg(m Message) {
	c.MsgChan <- m
}

func (c *Connection) Listen() error {
	for {
		// Read the message from the client
		_, message, err := c.WebConn.ReadMessage()
		if err != nil {
			log.Println("failed to read the message: ", err)
			return err
		}
		fmt.Printf("recv: %s\n", message)

		// append
		tmp := room.Get("username")
		tmp.AddMsg(Message{Type: "Server", Content: string(message)})
	}
}

func (c *Connection) SendMsg() error {
	for msg := range c.MsgChan {
		data, err := json.Marshal(msg)
		if err != nil {
			return errors.Wrap(err, "error in Marshal message")
		}

		// Send the message to the client
		if err = c.WebConn.WriteMessage(1, data); err != nil {
			return errors.Wrap(err, "failed to send the message")
		}
	}
	return nil
}

type Room struct {
	Room map[string]*Connection
	Lock sync.Mutex
}

func (r *Room) Get(id string) *Connection {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	return r.Room[id]
}

func (r *Room) Set(id string, conn *Connection) {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	r.Room[id] = conn
}

var (
	room = Room{Room: make(map[string]*Connection)}
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/connect", func(ctx *gin.Context) {
		log.Println("Connecting:", ctx.ClientIP())
		conn, err := NewConnection(ctx)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer conn.WebConn.Close()
		defer log.Println("Disconnection:", ctx.ClientIP())

		room.Set("username", conn)
		go func() {
			if err := conn.Listen(); err != nil {
				fmt.Println(err)
				close(conn.MsgChan)
			}
		}()

		if err := conn.SendMsg(); err != nil {
			fmt.Println(err)
		}
	})

	// css javascript
	router.Static("/static", "./static")

	// Start service
	router.Run(":5000")
}
