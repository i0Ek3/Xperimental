package asrt

import (
	"reflect"
	"testing"
)

func asrt(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
