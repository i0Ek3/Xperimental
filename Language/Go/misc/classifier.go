package main

import "fmt"

func main() {
	fmt.Println("classifier(1, -2, 13.0, 'hello', nil, true) will show you message like this: ")
	classifier(1, -2, 13.0, "hekko", nil, true)
}

// 类型判断：type-switch
//
// switch t := areaIntf.(type) {
// case *Square:
//	fmt.Printf("Type Square %T with value %v\n", t, t)
// case *Circle:
//	fmt.Printf("Type Circle %T with value %v\n", t, t)
// case nil:
//	fmt.Printf("nil value: nothing to check?\n")
// default:
//	fmt.Printf("Unexpected type %T\n", t)
// }

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case float32, float64:
			fmt.Printf("Param #%d is a float\n", i)
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		default:
			fmt.Printf("Param #%d is a unknow type\n", i)
		}
	}
}
