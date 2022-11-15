package main

import (
	"io"
)

// https://pkg.go.dev/bytes@go1.19.3

type readOp int8

const smallBufferSize = 64

type Buffer struct {
	// contents are the bytes buf[off : len(buf)]
	buf []byte
	// read at &buf[off], write at &buf[len(buf)]
	off int
	// last read operation, so that Unread* can work correctly.
	lastRead readOp
}

func NewBuffer(buf []byte) *Buffer {
	return &Buffer{buf: buf}
}

func NewBufferString(s string) *Buffer {
	return &Buffer{buf: []byte(s)}
}

func (b *Buffer) Bytes() []byte {
	return b.buf[b.off:]
}

func (b *Buffer) String() string {
	if b == nil {
		return "<nil>"
	}
	return string(b.buf[b.off:])
}

const opInvalid readOp = 0

func (b *Buffer) WriteString(s string) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(len(s))
	if !ok {
		m = b.grow(len(s))
	}
	return copy(b.buf[m:], s), nil
}

func (b *Buffer) tryGrowByReslice(n int) (int, bool) {
	if l := len(b.buf); n <= cap(b.buf)-l {
		b.buf = b.buf[:l+n]
		return l, true
	}
	return 0, false
}

// We just simplify the logic of grow()
func (b *Buffer) grow(n int) int {
	// grow grows the buffer to guarantee space for n more bytes.
	// It returns the index where bytes should be written.
	// If the buffer can't grow it will panic with ErrTooLarge.
	if n <= smallBufferSize && b.buf == nil {
		b.buf = make([]byte, n, smallBufferSize)
		return 0
	}
	return len(b.buf) - b.off
}

type Reader struct {
	s []byte
	// current reading index
	i int64
	// index of previous rune
	prevRune int
}

func NewReader(b []byte) *Reader {
	return &Reader{b, 0, -1}
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}
