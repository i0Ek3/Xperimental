package mswitch

import (
	"testing"
)

func TestSwitch(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		num := 100
		got := Switch(num)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
