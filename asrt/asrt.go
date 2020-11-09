package asrt

import (
	"reflect"
	"testing"
)

// Asrt compares the give value
func Asrt(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
