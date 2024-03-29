package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
}

type Req struct {
	A, B int
	Args int
}

type Resp struct {
	Result int
}

func (s *Server) Div(req Req, resp *Resp) error {
	// Simple validation
	if req.Args != 2 {
		log.Panic("wrong args")
	}
	if req.B == 0 {
		log.Panic("divider cannot be 0")
	}
	// Use Call will wait from here
	// time.Sleep(3 * time.Second)
	resp.Result = req.A / req.B
	return nil
}

func main() {
	// Register service first
	rpc.Register(new(Server))
	// Register handler
	rpc.HandleHTTP()
	// Then listen
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	// Run & Serve
	http.Serve(lis, nil)
}
