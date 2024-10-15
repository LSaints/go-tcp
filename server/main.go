package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
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
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestLines := strings.Split(string(buf), "\n")
	method := strings.Split(requestLines[0], " ")[0]
	requestPath := strings.Split(requestLines[0], " ")[1]
	fmt.Println(method, requestPath)

	err = validateMethod(method)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path, err := validateRequestPath(requestPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
	staticFile := getStaticContentFile(path)

	contentType := "html/text"
	if strings.HasSuffix(requestPath, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(requestPath, ".ico") {
		contentType = "image/x-icon"
	}
	fmt.Println(requestPath)
	response := returnHTTPResponse(staticFile, contentType)
	conn.Write([]byte(response))
}

func validateMethod(method string) error {
	if method == "GET" {
		fmt.Printf("metodo: %s suportado\n", method)
		return nil
	} else {
		return fmt.Errorf("metodo não permitido")
	}
}

func validateRequestPath(requestPath string) (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if requestPath == "/" || requestPath == "/index" {
		requestPath = "public/"
		staticFilePath := filepath.Join(rootDir, requestPath)
		return staticFilePath, nil
	}
	if requestPath == "/public/favicon.ico" {
		dir := filepath.Join(rootDir, "/public/favicon.ico")
		dir = strings.Replace(dir, "/index.html", "", 1)
		return dir, nil
	}
	return "", fmt.Errorf("caminho de solicitação inválido: %s", requestPath)
}

func getStaticContentFile(pathUri string) []byte {
	path := filepath.Join(pathUri, "index.html")
	if strings.Contains(pathUri, "favicon") {
		iconPath := strings.Replace(path, "/index.html", "", -1)
		path = iconPath
	}
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func returnHTTPResponse(content []byte, contentType string) string {
	httpResponse := fmt.Sprintf(`
		HTTP/1.1 200 OK\r\n
		Content-Type: %s\r\n
		Content-Length: %d\r\n\r\n

		%s
	`, contentType, len(content), content)

	return httpResponse
}
