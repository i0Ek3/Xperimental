package socket

import (
	"fmt"
	"io"
	"net"
)

var (
	host          = "www.baidu.com"
	port          = "80"
	remote        = host + ":" + port
	msg    string = "GET / \n"
	data          = make([]uint8, 4096)
	read          = true
	count         = 0
)

// open, write and read data from socket package
func Conn() bool {
	// create a socket
	con, err := net.Dial("tcp", remote)
	if err != nil {
		return false
	}

	// send a http GET request
	io.WriteString(con, msg)

	// read response from server
	for read {
		count, err = con.Read(data)
		read = (err == nil)
		fmt.Printf(string(data[0:count]))
	}
	con.Close()
	return true
}
