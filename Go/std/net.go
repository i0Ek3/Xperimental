package std

import (
	"context"
	"time"

	"golang.org/x/sync/singleflight"
)

// https://pkg.go.dev/net@go1.19.3

type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
	LocalAddr() Addr
	RemoteAddr() Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

type Addr interface {
	Network() string //  "tcp", "udp"
	String() string  // string form of address "192.0.2.1:25", "[2001:db8::1]:80"
}

type Dialer struct {
	Timeout       time.Duration
	Deadline      time.Time
	LocalAddr     Addr
	FallbackDelay time.Duration
	KeepAlive     time.Duration
	Resolver      *Resolver
}

type Resolver struct {
	Dial        func(ctx context.Context, network, address string) (Conn, error)
	lookupGroup singleflight.Group
}

func Dial(network, address string) (Conn, error) {
	var d Dialer
	return d.Dial(network, address)
}

func (d *Dialer) Dial(network, address string) (Conn, error) {
	return d.DialContext(context.Background(), network, address)
}

func (d *Dialer) DialContext(ctx context.Context, network, address string) (conn Conn, err error) {
	if ctx == nil {
		panic("nil context")
	}

	// ...

	return
}
