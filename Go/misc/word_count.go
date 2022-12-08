package main

import (
	"fmt"
)

func main() {
	str := " this is a  test with ver. 1.0. "
	fmt.Println(wordCount(str))
}

// O(n)/O(n)
func wordCount(str string) int {
	idx := -1
	res := []string{}
	for i, v := range str {
		if !isSep(v) {
			if idx == -1 {
				idx = i
			}
			continue
		} else if isDecimal(v, str, i-1, i+1) {
			continue
		} else {
			if idx != -1 {
				res = append(res, str[idx:i])
				idx = -1
			}
		}
	}

	if idx != -1 {
		res = append(res, str[idx:])
	}

	return len(res)
}

func isSep(r rune) bool {
	for _, v := range []rune{' ', '.'} {
		if v == r {
			return true
		}
	}
	return false
}

func isN(i rune) bool {
	return i >= '0' && i <= '9'
}

func isDecimal(r rune, s string, i, j int) bool {
	if r == '.' && isN(b2r(s[i])) && isN(b2r(s[j])) {
		return true
	}
	return false
}

func b2r(b byte) rune {
	return []rune(string(b))[0]
}

func r2b(r rune) byte {
	return []byte(string(r))[0]
}
