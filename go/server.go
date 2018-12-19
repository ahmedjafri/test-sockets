package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(c)
	}
}

func handleRequest(c net.Conn) {
	fmt.Printf("Connected to client!\n")
	for {
		b := make([]byte, 100)
		_, err := c.Read(b)
		if err != nil {
			panic(err)
		}
	}
}
