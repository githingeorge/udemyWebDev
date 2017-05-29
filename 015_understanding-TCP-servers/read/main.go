package main

import (
	"bufio"
	"fmt"
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
			continue
		}

		go handle(conn)

	}
}
func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("code got here")

}
