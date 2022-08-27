package main

import (
    "fmt"
    "time"
)

type Pool struct {
    Work chan func()
    Num  chan struct{}
}

func New(size int) *Pool {
    return &Pool{
        Work: make(chan func()),
        Num:  make(chan struct{}, size),
    }
}

func (p *Pool) NewTask(task func()) {
    select {
    case p.Work <- task:
    case p.Num <- struct{}{}:
        go p.worker(task)
    //default:
    }
}

func (p *Pool) worker(task func()) {
    defer func() {
        <-p.Num
    }()

    for {
        task()
        task = <-p.Work
    }
}

func main() {
    pool := New(2)
    for i := 1; i < 5; i++ {
        pool.NewTask(func(){
            time.Sleep(2 * time.Second)
            fmt.Println("task running...", time.Now())
        })
    }
    time.Sleep(5 * time.Second)
}
