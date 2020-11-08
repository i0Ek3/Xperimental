package mxml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestMyXML(t *testing.T) {
	input := "<Person><FirstName>Xxx</FirstName><LastName>Xxx</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)

	got := mxml(p)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
