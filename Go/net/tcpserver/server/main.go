package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8888")
	lis, _ := net.ListenTCP("tcp", addr)

	for {
		conn, err := lis.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handle(conn)
	}
}

func handle(conn *net.TCPConn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String() + " connected: " + string(buf[0:n]))
		msg := "got it: " + string(buf[0:n])
		conn.Write([]byte(msg))
	}
}
