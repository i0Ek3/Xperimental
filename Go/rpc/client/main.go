package main

import (
	"fmt"
	"net/rpc"
	"time"
)

type Req struct {
	A, B int
	Args int
}

type Resp struct {
	Result int
}

func main() {
	req := Req{A: 1, B: 2, Args: 2}
	var resp Resp

	cli, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	// Call will wait until server finished out
	// _ = cli.Call("Server.Add", req, &resp)

	ca := cli.Go("Server.Add", req, &resp, nil)

	for {
		select {
		case <-ca.Done:
			fmt.Println(resp)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("I'm here waitting for you.")
		}
	}
}
