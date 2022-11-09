package main

import (
	"context"
	"fmt"
	"time"
)

func monitor(ctx context.Context, num int) {
	for {
		select {
		case v := <-ctx.Done():
			fmt.Printf("Monitor %v recv %v.\n", num, v)
			return
		default:
			fmt.Printf("Monitor %v monitoring...\n", num)
			time.Sleep(2 * time.Second)
		}
	}
}

func ctxTest() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 5; i++ {
		go monitor(ctx, i)
	}

	time.Sleep(time.Second)
	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("Main program exited!")
}

func ctxTest2() {
	ctx := context.Background()
	pass(ctx)

	ctx = context.WithValue(ctx, "passID", "i0Ek3")
	pass(ctx)
}

func pass(ctx context.Context) {
	passID, ok := ctx.Value("passID").(string)
	if ok {
		fmt.Printf("passID = %s\n", passID)
	} else {
		fmt.Println("passID is null")
	}
}

func ctxTest3() {
	ctx := context.WithValue(context.Background(), "i0Ek3", 18)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	ch1 := make(chan struct{})
	ch2 := make(chan int)
	go wait(ctx, ch1, ch2)
	for i := 1; i <= 10; i++ {
		ch2 <- i
	}
	ch1 <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("Test3 bye")
}

func wait(ctx context.Context, ch1 chan struct{}, ch2 chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case c2 := <-ch2:
			fmt.Printf("c2 = %d\n", c2)
		case <-ctx.Done():
			fmt.Println(ctx.Value("i0Ek3"))
			return
		}
	}
}

func main() {
	ctxTest3()
}
