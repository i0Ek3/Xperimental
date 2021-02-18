package main

import (
	"fmt"
	"time"
)

// Timer为定时器，Ticker为计时器
//
// time.Ticker对象以指定的时间间隔重复的向通道 C 发送时间值
// type Ticker struct {
//         C <- chan Time
//         ....
// }
//
// 结果通道是必须要带缓冲的

func main() {
	tick := time.Tick(1e8)
	boom := time.After(5e8) // 只发送一次时间
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
