package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Erro ao enviar:", err)
			return
		}
		fmt.Printf("menssagem enviada > %s\n", conn.RemoteAddr().String())
	}
}
