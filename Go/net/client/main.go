package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8888")
	conn, _ := net.DialTCP("tcp", nil, addr)
	reader := bufio.NewReader(os.Stdin)
	for {
		byt, _, _ := reader.ReadLine()
		conn.Write(byt)
		feedback := make([]byte, 1024)
		n, _ := conn.Read(feedback)
		fmt.Println(feedback[0:n])
	}
}
