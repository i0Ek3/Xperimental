package main

import (
	"fmt"
	"sort"
)

var (
	seq = map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
		"d": 4,
		"f": 6,
		"e": 5}
)

func main() {
	sorted()
	inverted()
	test1()
	test2()
}

func sorted() {
	fmt.Println("Original Seq:")
	for k, v := range seq {
		fmt.Printf("Key = %v, Val = %v\n", k, v)
	}
	keys := make([]string, len(seq))
	i := 0
	for k := range seq {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("Sorted Seq:")
	for _, k := range keys {
		fmt.Printf("Key = %v, Val = %v\n", k, seq[k])
	}
}

func inverted() {
	invertMap := make(map[int]string, len(seq))
	for k, v := range seq {
		invertMap[v] = k
	}
	fmt.Println("Inverted:")
	for k, v := range invertMap {
		fmt.Printf("Key = %v, Val = %v\n", k, v)
	}

}

func test1() {
	var mapStr1 map[string]int
	var mapStr2 map[string]int

	mapStr1 = map[string]int{"a": 1, "b": 2, "c": 3}
	mapCreated := make(map[string]float32)
	mapStr2 = mapStr1

	mapCreated["key1"] = 1.0
	mapCreated["key2"] = 2.0
	mapStr2["two"] = 3

	fmt.Printf("mapStr1 at \"one\" is: %d\n", mapStr1["one"])
	fmt.Printf("mapCreated at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("mapStr2 at \"two\" is: %d\n", mapStr1["two"])
	fmt.Printf("mapStr1 at \"ten\" is: %d\n", mapStr1["ten"])
}

func test2() {
	map_func := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Println(map_func)
	//map_slice := make(map[int][]int)
}
