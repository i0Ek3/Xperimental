package main

import (
	"fmt"
	"sync"
)

func Common() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func WithDo() {
	one := &sync.Once{}

	for i := 0; i < 5; i++ {
		one.Do(func() {
			fmt.Println(i)
		})
	}
}

func main() {
	fmt.Println("------WithOnce-------")
	WithDo()
	fmt.Println("-----WithCommon-------")
	Common()
}
