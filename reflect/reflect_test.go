package main

import (
	"reflect"
	"testing"

	a "github.com/i0Ek3/asrt"
)

func TestReflect(t *testing.T) {
	x := "reflect"

	t.Run("valueof", func(t *testing.T) {
		got := valueof(x)
		want := "reflect"

		a.Asrt(t, got, want)
	})

	t.Run("typeof", func(t *testing.T) {
		got := typeof(x)
		want := reflect.String

		a.Asrt(t, got, want)
	})
}
