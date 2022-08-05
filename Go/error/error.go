package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("not found")

func main() {
	fmt.Printf("error: %v", errNotFound)
}
