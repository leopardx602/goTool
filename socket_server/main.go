package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	fmt.Println("Connected!")
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maxium request length to 128B to prevent flood attack

	connStatus := true

	// send message
	go func() {
		for connStatus {
			time.Sleep(2 * time.Second)
			daytime := time.Now().String()
			fmt.Println("server:", daytime)
			conn.Write([]byte(daytime))
		}
	}()

	// receive message
	for {
		read_len, err := conn.Read(request) // block
		if err != nil || read_len == 0 {
			fmt.Println("Disconnected!")
			connStatus = false
			break // connection already closed by client
		} else {
			fmt.Println(strings.TrimSpace(string(request[:read_len])))
		}
		request = make([]byte, 128) // clear last read content
	}
}

func main() {
	service := ":5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
