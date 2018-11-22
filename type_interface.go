package main

import (
        "fmt"
        "math"
)

type Square struct {
        side float32 
}

type Circle struct {
        radius float32 
}

type Shaper interface {
        Area() float32 
}

func main() {
        var areaIntf Shaper // 定义了一个接口变量
        sq := new(Square)
        sq.side = 5

        areaIntf = sq

        // 类型断言
        if t, ok := areaIntf.(*Square); ok {
                fmt.Printf("The type of areaIntf is: %T\n", t)
        }

        if u, ok := areaIntf.(*Circle); ok {
                fmt.Printf("The type of areaIntf is: %T\n", u)
        } else {
                fmt.Println("areaIntf doesn't contain a variable of type Circle")                
        }
}

func (sq *Square) Area() float32 {
        return sq.side * sq.side 
}

func (ci *Circle) Area() float32 {
        return ci.radius * ci.radius * math.Pi
}


