package mxml

import (
	"encoding/xml"
	"fmt"
)

var (
	t     xml.Token
	token xml.Token
)

func mxml(p *xml.Decoder) bool {
	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
			return true
		case xml.EndElement:
			fmt.Println("End of token.")
			return true
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
			return true
		default:
			return false
		}
	}
	return false
}
