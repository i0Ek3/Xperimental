package std

import (
	"io"
	"sync"
)

type Logger struct {
	mu        sync.Mutex // ensures atomic writes; protects the following fields
	prefix    string     // prefix on each line to identify the logger (but see Lmsgprefix)
	flag      int        // properties
	out       io.Writer  // destination for output
	buf       []byte     // for accumulating text to write
	isDiscard int32      // atomic boolean: whether out == io.Discard
}

func New(out io.Writer, prefix string, flag int) *Logger {
	l := &Logger{out: out, prefix: prefix, flag: flag}
	if out == io.Discard {
		l.isDiscard = 1
	}
	return l
}
