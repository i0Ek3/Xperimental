package std

import (
	"errors"
	"io"
	"log"
	"net"
	"sync"
)

type ServerError string

func (e ServerError) Error() string {
	return string(e)
}

var ErrShutdown = errors.New("connection is shut down")

// Call represents an active RPC.
type Call struct {
	ServiceMethod string
	Args          any
	Reply         any
	Error         error
	Done          chan *Call
}

type Client struct {
	codec ClientCodec

	reqMutex sync.Mutex // protects following
	request  Request

	mutex    sync.Mutex // protects following
	seq      uint64
	pending  map[uint64]*Call
	closing  bool // user has called Close
	shutdown bool // server has told us to stop
}

type ClientCodec interface {
	WriteRequest(*Request, any) error
	ReadResponseHeader(*Response) error
	ReadResponseBody(any) error

	Close() error
}

type Request struct {
	ServiceMethod string   // format: "Service.Method"
	Seq           uint64   // sequence number chosen by client
	next          *Request // for free list in Server
}

type Response struct {
	ServiceMethod string    // echoes that of the Request
	Seq           uint64    // echoes that of the request
	Error         string    // error, if any.
	next          *Response // for free list in Server
}

type Server struct {
	serviceMap sync.Map   // map[string]*service
	reqLock    sync.Mutex // protects freeReq
	freeReq    *Request
	respLock   sync.Mutex // protects freeResp
	freeResp   *Response
}

func NewServer() *Server {
	return &Server{}
}

// DefaultServer is the default instance of *Server.
var DefaultServer = NewServer()

func Dial(network, address string) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewClient(conn), nil
}

func NewClient(conn io.ReadWriteCloser) *Client {
	return &Client{}
}

// asynchronously
func (c *Client) Go(method string, args, reply any, done chan *Call) *Call {
	call := new(Call)
	call.ServiceMethod = method
	call.Args = args
	call.Reply = reply
	if done == nil {
		done = make(chan *Call, 10)
	} else {
		if cap(done) == 0 {
			log.Panic("chan is unbuffered")
		}
	}
	call.Done = done
	c.send(call)
	return call
}

func (c *Client) send(call *Call) {
	c.reqMutex.Lock()
	defer c.reqMutex.Unlock()

	// register call
	c.mutex.Lock()
	if c.shutdown || c.closing {
		c.mutex.Unlock()
		call.Error = ErrShutdown
		call.done()
		return
	}
	seq := c.seq
	c.seq++
	c.pending[seq] = call
	c.mutex.Unlock()

	// encode and send req
	c.request.Seq = seq
	c.request.ServiceMethod = call.ServiceMethod
	// write req and handle error
}

func (c *Call) done() {
	select {
	case c.Done <- c:
	default:
		if false {
			log.Println("discarding reply cause of insufficient Done chan cap")
		}
	}
}

func (c *Client) Call(method string, args, reply any) error {
	call := <-c.Go(method, args, reply, make(chan *Call, 1)).Done
	return call.Error
}
