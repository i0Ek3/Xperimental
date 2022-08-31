package main

import "syscall/js"

func main() {
	str := "Hi"
	alert := js.Global().Get("alert")
	alert.Invoke(str)
}
