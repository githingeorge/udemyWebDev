package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	ls, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer ls.Close()

	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "hii  this is githin")
}
