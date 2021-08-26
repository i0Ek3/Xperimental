package main

import (
	"fmt"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for i := 0; i < data.Len(); i++ {
		for j := 0; j < data.Len()-i; j++ {
			if data.Less(i+1, i) {
				data.Swap(i+1, i)
			}
		}
	}
}

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42}
	a := IntArray(data) //conversion to type IntArray
	Sort(a)
	fmt.Printf("The sorted array is: %v\n", a)
}
