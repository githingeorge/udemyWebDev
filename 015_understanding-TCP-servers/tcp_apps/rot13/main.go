package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	defer conn.Close()

	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		rbs := rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n\n", ln, rbs)
	}
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		// asscii 97 -122
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}

	return r13
}
