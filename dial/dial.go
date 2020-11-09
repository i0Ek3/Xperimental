package dial

import (
	"fmt"
	"net"
	"os"
)

// Dial checks if connection is ok
func Dial() bool {
	conn, err := net.Dial("tcp", "220.181.38.149:80")
	return checkConnection(conn, err)
	//conn, err = net.Dial("udp", "ipv4:port")
	//checkConnection(conn, err)
	//conn, err = net.Dial("tcp", "ipv6:port")
	//checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) bool {
	if err != nil {
		fmt.Printf("error %v connecting!", err)
		return false
		os.Exit(1)
	}
	fmt.Printf("Connection is made with %v\n", conn)
	return true
}
