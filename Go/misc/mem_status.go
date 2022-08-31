package main

import (
        "fmt"
        "runtime"
)

// 可使用以下语句来实现在将对象从内存中移除时做一些特殊的操作
// runtime.SetFinalizer(obj, func(obj *typeObj))

func main() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        fmt.Printf("%d Kb\n", m.Alloc / 1024)
}
