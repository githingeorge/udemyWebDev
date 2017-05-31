package main

import (
	"bufio"
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
	// err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}

	}
	fmt.Println("code got here")

	io.WriteString(conn, "I see you connected")
}
