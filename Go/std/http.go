package std

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"mime/multipart"
	"net"
	"net/url"
	"time"
)

// https://pkg.go.dev/net/http@go1.19.3

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type Header map[string][]string

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

type Request struct {
	Method           string
	URL              *url.URL
	Proto            string // "HTTP/1.0"
	ProtoMajor       int    // 1
	ProtoMinor       int    // 0
	Header           Header
	Body             io.ReadCloser
	GetBody          func() (io.ReadCloser, error)
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Host             string
	Form             url.Values
	PostForm         url.Values
	MultipartForm    *multipart.Form
	RemoteAddr       string
	RequestURI       string
	TLS              *tls.ConnectionState
	Response         *Response
	ctx              context.Context
}

type Response struct {
	Status           string // e.g. "200 OK"
	StatusCode       int    // e.g. 200
	Proto            string // e.g. "HTTP/1.0"
	ProtoMajor       int    // e.g. 1
	ProtoMinor       int    // e.g. 0
	Header           Header
	Body             io.ReadCloser
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Uncompressed     bool
	Request          *Request
	TLS              *tls.ConnectionState
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

type ConnState int

type Server struct {
	Addr              string
	Handler           Handler
	TLSConfig         *tls.Config
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
	TLSNextProto      map[string]func(*Server, *tls.Conn, Handler)
	ConnState         func(net.Conn, ConnState)
	ErrorLog          *log.Logger
	BaseContext       func(net.Listener) context.Context
	ConnContext       func(ctx context.Context, c net.Conn) context.Context
}

// these functions kinda complicated, so we just return the nil
func (srv *Server) Close() error                                            { return nil }
func (srv *Server) ListenAndServe() error                                   { return nil }
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error        { return nil }
func (srv *Server) Serve(l net.Listener) error                              { return nil }
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error { return nil }
func (srv *Server) Shutdown(ctx context.Context) error                      { return nil }
