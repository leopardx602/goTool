package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func checkError(err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err)
		time.Sleep(2 * time.Second)
	}
	return err
}

func sendLoop(conn net.Conn, connStatus *bool) {
	for *connStatus {
		_, err := conn.Write([]byte("hello server"))
		checkError(err)
		time.Sleep(2 * time.Second)
	}
}

func receiveLoop(conn net.Conn, connStatus *bool) {
	request := make([]byte, 128)

	for {
		data, err := conn.Read(request)
		if err != nil || data == 0 {
			fmt.Println("Disconnected!")
			*connStatus = false
			break // connection already closed by client
		} else {
			fmt.Println(strings.TrimSpace(string(request[:data])))
		}
		request = make([]byte, 128) // clear last read content
	}

}

func connectToServer() {
	service := "127.0.0.1:5000"
	connStatus := false

	for {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
		if err != nil {
			fmt.Println("error:", err)
			time.Sleep(2 * time.Second)
			continue
		}
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Println("error:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		connStatus = true
		go receiveLoop(conn, &connStatus)
		sendLoop(conn, &connStatus)
	}
}

func main() {
	go connectToServer()

	var msg string
	for {
		fmt.Print("please input: ")
		fmt.Scanln(&msg)
		fmt.Println("your input is:", msg)

		if msg == "exit" {
			break
		}
	}
}
