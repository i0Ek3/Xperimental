package socket

import (
	"testing"

	a "asrt"
)

func TestConn(t *testing.T) {
	got := Conn()
	want := true

	a.Asrt(t, got, want)
}
