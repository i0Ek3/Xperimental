package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial("ws://localhost:8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
	go send(conn)

	for {
		m, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(m, string(p))
	}
}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, line)
	}
}
