package main

import (
	"fmt"
	"os"
)

var err error

func init() {
	if err != nil {
		panic("ERROR occurred:" + err.Error())
	}
}

var user = os.Getenv("USER")

func check() {
	if user == "" {
		panic("Unknown user: no value or $USER")
	}
}

func main() {
	check()
	fmt.Println("Starting...")
	panic("A server error occurred: Stop it now!")
	//fmt.Println("Ending...")
}
