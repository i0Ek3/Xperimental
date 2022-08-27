package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(<-chan string)
	ch2 := make(<-chan string)

	go priority(ch1, ch2)
	time.Sleep(time.Second)
}

func priority(ch1, ch2 <-chan string) {
	for {
		select {
		case val1 := <-ch1:
            fmt.Println("val1 printed:", val1)
		case val2 := <-ch2:
		priority:
			for {
				select {
				case val1 := <-ch1:
                    fmt.Println("val1 printed but with round 2:", val1)
				default:
					break priority
				}
			}
            fmt.Println("val2 printed:", val2)
			//default:
		}
	}
}
