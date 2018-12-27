<<<<<<< HEAD
package main
=======
package parse
>>>>>>> 777c5bda2501a2e54ee844b2181a09c4b2b19e4c

import (
        "fmt"
        "strings"
        "strconv"
)

<<<<<<< HEAD
//
// 自定义包实现者应该遵守的最佳实践:
//      在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()
//      向包的调用者返回错误值（而不是 panic）
//

=======
>>>>>>> 777c5bda2501a2e54ee844b2181a09c4b2b19e4c
type ParseError struct {
        Index   int
        Word string 
        Err   error
}

func (e *ParseError) String() string {
        return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

func Parse(input string) (numbers []int, err error) {
        defer func() {
                if r := recover(); r != nil {
                        var ok bool
                        err, ok = r.(error)
                        if !ok {
                                err = fmt.Errorf("pkg: %v", r)
                        }
                }
        }()
        fields := strings.Fields(input)
        numbers = fields2numbers(fields)
        return 
}

func fields2numbers(fields []string) (numbers []int) {
        if len(fields) == 0 {
                panic("No words to parse.")    
        }
        for idx, field := range fields {
                num, err := strconv.Atoi(field)
                if err != nil {
                        panic(&ParseError{idx, field, err})
                }
                numbers = append(numbers, num)
        }
        return 
}
<<<<<<< HEAD

func main() {
        var examples = []string{
                "1 2 3 4 5",
                "100 50 25 12.5 6.25",
                "2 2 = 4",
                "1st class",
                " ",
        }

        for _, ex := range examples {
                fmt.Printf("Parsing %q: \n", ex)
                nums, err := Parse(ex)
                if err != nil {
                        fmt.Println(err)
                        continue
                }
                fmt.Println(nums)
        }
}
=======
>>>>>>> 777c5bda2501a2e54ee844b2181a09c4b2b19e4c
