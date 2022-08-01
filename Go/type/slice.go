package main

import "fmt"

func main() {
    testCopy()
}

func test() {
	var arr = [5]int{1, 2, 3, 4, 5}
	sum(arr[:])
	fmt.Println("-------------------------\n")
	printSlice()
	fmt.Println("-------------------------\n")
	makeSlice()
}

// slice as paramete
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func printSlice() {
	var arr1 [6]int
	var slice1 []int = arr1[1:4]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("len(arr1) = %d\n", len(arr1))
	fmt.Printf("len(slice1) = %d\n", len(slice1))
	fmt.Printf("cap(slice1) = %d\n", cap(slice1))

	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("len(slice1) = %d\n", len(slice1))
	fmt.Printf("cap(slice1) = %d\n", cap(slice1))

}

func makeSlice() {
	// func make([]T, len, cap)
	var slice1 []int = make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\nlen(slice1) = %d\n", len(slice1))
	fmt.Printf("\ncap(slice1) = %d\n", cap(slice1))
}

func testCopy() {
    s1 := []string{"1", "2", "3"}
    fmt.Printf("s1 = %v, &s1 = %p\n", s1, &s1)

    s2 := s1
    fmt.Printf("s2 = %v, &s2 = %p\n", s2, &s2)

    s2[0] = "5"
    fmt.Printf("s1 = %v, &s1 = %p\n", s1, &s1)
    fmt.Printf("s2 = %v, &s2 = %p\n", s2, &s2)

    s1 = append(s1, "4")
    fmt.Printf("s1 = %v, &s1 = %p\n", s1, &s1)
    fmt.Printf("s2 = %v, &s2 = %p\n", s2, &s2)
}
