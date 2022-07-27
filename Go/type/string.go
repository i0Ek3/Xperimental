package mstring

import (
	"strings"
)

func mstring(oldString, newString string) bool {
	if !strings.Contains(oldString, newString) {
		if strings.ToUpper(oldString) == newString {
			return true
		}
	}
	return false
}
