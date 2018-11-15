package main 
import "fmt"

func main() {
        seasons := []string{"Spr", "Sum", "Aut", "Win"}
        for index, season := range seasons {
                fmt.Printf("Season %d = %s\n", index, season)
        }

        var season string
        for _, season = range seasons {
                fmt.Printf("%s\n", season)
        }
}
