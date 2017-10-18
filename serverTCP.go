package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("use telnet localhost 8000")
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err) // Stop the runtime if fail
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // For example, aborted connection
			continue
		}
		go manageWith(conn) // We handle one connection at time
	}
}

var id = 0

func manageWith(c net.Conn) {
	id++
	fmt.Println("Connections: ", id)
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return // For example, if the client disconnect
		}
		time.Sleep(1 * time.Second)
	}
}
