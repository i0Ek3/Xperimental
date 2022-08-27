package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
    res := make([]K, 0, len(m))
    for k := range m {
        res = append(res, k)
    }
    return res
}

type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (l *List[T]) Push(v T) {
    if l.tail == nil {
        l.head = &element[T]{val: v}
        l.tail = l.head
    } else {
        l.tail.next = &element[T]{val: v}
        l.tail = l.tail.next
    }
}

func (l *List[T]) GetAll() []T {
    var elems []T
    for e := l.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
    var m = map[int]string{1: "2", 2: "4", 4: "8"}
    fmt.Println("keys m:", MapKeys(m))

    fmt.Println(MapKeys[int, string](m))

    l := List[int]{}
    l.Push(10)
    l.Push(13)
    l.Push(23)
    fmt.Println("list:", l.GetAll())
}
