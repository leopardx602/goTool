package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Discount bool   `json:"discount"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.GET("/connect", func(ctx *gin.Context) {
		log.Println("Connecting:", ctx.ClientIP())
		var upgrader = websocket.Upgrader{}
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println("Failed to Upgrade: ", err)
			return
		}
		defer conn.Close()
		defer log.Println("Disconnection:", ctx.ClientIP())
		for {
			// Read the message from the client
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Failed to read the message: ", err)
				break
			}
			fmt.Printf("recv: %s\n", message)

			// Send the message to the client
			if err = conn.WriteMessage(mt, message); err != nil {
				log.Println("Failed to send the message: ", err)
				break
			}
		}
	})

	// css javascript
	router.Static("/static", "./static")

	// Start service
	router.Run(":5000")
}
