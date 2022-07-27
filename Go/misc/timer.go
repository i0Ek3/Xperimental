package main

import (
	"fmt"
	"time"
)

// Time shows tick and boom time
func Time(tick time.Tick, boom time.Duration) {
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("   .")
			time.Sleep(5e7)
		}
	}
}
