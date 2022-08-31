package main

import (
    "fmt"
    "sync"
    "time"
)

func add1(cnt *int, wg *sync.WaitGroup) {
    for i := 0; i < 1000; i++ {
        *cnt += 1
    }
    wg.Done()
}

func add2(cnt *int, wg *sync.WaitGroup, lock *sync.Mutex) {
    for i := 0; i < 1000; i++ {
        lock.Lock()
        *cnt += 1
        lock.Unlock()
    }
    wg.Done()
}

func check() {
    var wg sync.WaitGroup
    lock := &sync.Mutex{}
    
    cnt := 0
    wg.Add(3)
    
    go add2(&cnt, &wg, lock)
    go add2(&cnt, &wg, lock)
    go add2(&cnt, &wg, lock)

    wg.Wait()
    fmt.Println("cnt: ", cnt)
}

func check2() {
    l := &sync.RWMutex{}
    l.Lock()

    for i := 0; i < 5; i++ {
        go func(i int) {
            fmt.Printf("goroutine %d runing...\n", i)
            l.RLock()
            fmt.Printf("goroutine %d get lock, sleep and release lock\n", i)
            time.Sleep(time.Second)
            l.RUnlock()
        }(i)
    }
    time.Sleep(2*time.Second)
    fmt.Println("Releasing lock...")
    l.Unlock()

    l.Lock()
    fmt.Println("exiting...")
    l.Unlock()
}

func main() {
   check2() 
}
