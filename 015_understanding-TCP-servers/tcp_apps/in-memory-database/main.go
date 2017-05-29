package main

import (
	"bufio"
	"fmt"
	"io"
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
	// instructions
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	database := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	defer conn.Close()

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		if len(fs) < 2 {
			fmt.Fprintln(conn, "No data")
			continue
		}
		k := fs[1]
		switch fs[0] {
		case "GET":
			if v, ok := database[k]; ok {
				fmt.Fprintf(conn, "%s - %s", k, v)
			} else {
				fmt.Fprintln(conn, "No data")
			}
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "expected a value")
			}
			v := fs[2]
			database[k] = v
		case "DEL":
			delete(database, k)
		default:
			fmt.Fprintln(conn, "Invalid command")
			continue
		}
		fmt.Fprintln(conn, "")

	}
}
