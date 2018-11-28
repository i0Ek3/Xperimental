package main

import (
        //"fmt"
        "os"
        "strconv"
)

type Stringer interface {
        String() string 
}

type Temprature float64 

func (t Temprature) String() string {
        return strconv.FormatFloat(float64(t), 'f', 1, 64) + " °C"
}

type Day int
type Temp int

var dayName = []string{"Monday", "Tuesday", "Wednesay", "Thursday", "Friday", "Saturday", "Sunday"}
var tempValue = []string{"0.0", "15.7", "25.5", "-10.1", "-40.0", "37.7", "1.7"}

func (d Day) String() string {
        return dayName[d]
}

func (t Temp) String() string {
        return tempValue[t]
}

func print(args ...interface{}) {
        for i, arg := range args {
                if i > 0 {
                        os.Stdout.WriteString(" ")
                }   
                switch a := arg.(type) {
                        case Stringer: os.Stdout.WriteString(a.String())
                        case int: os.Stdout.WriteString(strconv.Itoa(a))
                        case string: os.Stdout.WriteString(a)
                        default: os.Stdout.WriteString("Unknown type!")
                }
        }
}

func main() {
        print(Day(3), "was", Temprature(35.8))
}

