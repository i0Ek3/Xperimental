package mswitch

import (
	"fmt"
)

// Switch
func Switch(num int) bool {
	return func(i int) bool {
		switch {
		case i > 10:
			fmt.Println("i > 10")
		default:
		}
		return true
	}(num)
}
