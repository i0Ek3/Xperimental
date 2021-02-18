package main

import (
	"fmt"
	"math"
)

// Square
type Square struct {
	side float32
}

// Area
func (s *Square) Area() float32 {
	return s.side * s.side
}

// Circle
type Circle struct {
	radius float32
}

// Area
func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

// Shaper
type Shaper interface {
	Area() float32
}

func main() {
	var areaIntfs Shaper // 定义了一个接口变量
	var areaIntfc Shaper // 定义了一个接口变量

	sq := new(Square)
	sq.side = 5.0

	areaIntfs = sq

	ci := new(Circle)
	ci.radius = 2.0
	areaIntfc = ci

	// 类型断言
	if t, ok := areaIntfs.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	if u, ok := areaIntfc.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf doesn't contain a variable of type Circle")
	}
}
