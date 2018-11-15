package main

import "fmt"

func main() {

//LABEL1:
        var words string = "I love singing."
        for i := 0; i < len(words); i++ {
                fmt.Printf(words[i])
        }

        for i := 0; i < 5; i++ {
                fmt.Printf("This is the %d iteration\n", i)
        }


        for i := 0; i < 5; i++ {
                for j := 1; j < i; j++ {
                        fmt.Println(j)    
                }
        }

        var h int = 3
        for h < 5 {
                h++
                fmt.Printf("Te far be lo i is now : %d\n", h)
        }

        for i := 0; i < 3; i++ {
                for j := 1; j < 10; j++ {
                        if j == 6 {
                                continue //LABEL1
                        } else if j > 6 {
                                break
                        }  
                        print(j)
                }
                print(i)
                print(" ")
        }
}


