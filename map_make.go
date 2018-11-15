package main
import "fmt"

func main() {
        var mapStr1 map[string] int
        var mapStr2 map[string] int

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
