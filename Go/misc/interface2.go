package main

import "fmt"

type X interface {
	Fire(string) string
	Blood() int
}

type Alterman struct {
	Status string
	blood  int
}

func (a Alterman) Fire(status string) string {
	return "Alterman on the fire!"
}

func (a Alterman) Blood() int {
	return a.blood
}

type Spiderman struct {
	Status string
	blood  int
}

func (s Spiderman) Fire(status string) string {
	return "Spiderman on the fire!"
}

func (s Spiderman) Blood() int {
	return s.blood
}

var Everything X

func main() {
	a := Alterman{
		Status: "fire",
		blood:  70,
	}
	Link(a)
	fmt.Println(Everything.Fire(a.Status))
}

func Link(x X) {
	Everything = x
}
