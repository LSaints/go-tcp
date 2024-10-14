package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		requestLines := strings.Split(string(buf), "\n")
		method := strings.Split(requestLines[0], " ")[0]
		path := strings.Split(requestLines[0], " ")[1]
		fmt.Println(method, path)

	}
}
