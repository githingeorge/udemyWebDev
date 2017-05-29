package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\n Hello from TCP server")
		fmt.Fprintln(conn, "how is your day?")
		fmt.Fprintf(conn, "%v", "well i hope")

		conn.Close()
	}
}
