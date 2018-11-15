package main

import (
    "fmt"
    "runtime"
)

//
// if nil != 0 {
//    // fuck off
// } else if nil == 0 {
//   // get out of here
// } else { // these braces and else must be in one line
//   // well done
// }
//
//
//

func ex(x, y int) int {
    if x > y {
        y = x
        return y
    }
    return x
}

func main() {
    if runtime.GOOS == "windows" {
        fmt.Printf("motherfucker,,,,,,\n")
    } else if runtime.GOOS == "darwin" {
        fmt.Printf("good choice!!!!!!!\n")
    } else {
        fmt.Printf("not bad.\n")
    }
    

    if ex(8,7) > 5 {
        fmt.Printf("ex(8, 7) large than 5!\n")
    } else {
        fmt.Printf("ex(8, 7) not large than 5!\n")
    }
}





 


