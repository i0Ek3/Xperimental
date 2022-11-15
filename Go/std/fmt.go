package std

import (
	"io"
	"os"
)

// https://pkg.go.dev/fmt@go1.19.3

type Stringer interface {
	String() string
}

type Formatter interface {
	Format(f State, verb rune)
}

type State interface {
	Write(b []byte) (n int, err error)
	Width() (wid int, ok bool)
	Precision() (prec int, ok bool)
	Flag(c int) bool
}

func Errorf(format string, a ...any) error { return nil }

func Print(a ...any) (n int, err error)                 { return Fprint(os.Stdout, a...) }
func Println(a ...any) (n int, err error)               { return }
func Printf(format string, a ...any) (n int, err error) { return }

func Fprintf(w io.Writer, format string, a ...any) (n int, err error) { return }
func Fprint(w io.Writer, a ...any) (n int, err error)                 { return }
func Fprintln(w io.Writer, a ...any) (n int, err error)               { return }

func Sprint(a ...any) string                 { return "" }
func Sprintf(format string, a ...any) string { return "" }
func Sprintln(a ...any) string               { return "" }
