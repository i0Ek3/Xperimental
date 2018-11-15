package main
import "fmt"

func main() {
        map_func := map[int]func() int {
                1: func() int { return 10 },  
                2: func() int { return 20 },  
                3: func() int { return 30 },  
        }
        fmt.Println(map_func)

        //map_slice := make(map[int][]int)
                
}
