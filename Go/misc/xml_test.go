package misc

import (
	"encoding/xml"
	"strings"
	"testing"

	a "asrt"
)

func TestMyXML(t *testing.T) {
	input := "<Person><FirstName>Xxx</FirstName><LastName>Xxx</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)

	got := mxml(p)
	want := true

	a.Asrt(t, got, want)
}
