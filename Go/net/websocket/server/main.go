package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, req *http.Request) {
	var conns []*websocket.Conn
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	conns = append(conns, conn)
	for {
		m, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		for i := range conns {
			conns[i].WriteMessage(websocket.TextMessage, []byte("->"+string(p)+"?"))
		}
		fmt.Println(m, string(p))
	}
	log.Println("closed")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
