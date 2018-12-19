package main

import (
	"net"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	b := make([]byte, 100)
	for {
		c.Read(b)
	}
}
