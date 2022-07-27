package main

import (
        "fmt"
        "math"
        "math/big"
)

// 大整数运算

func main() {
        ia := big.NewInt(math.MaxInt64)
        ib := ia
        ic := big.NewInt(1994)
        id := big.NewInt(1)
        id.Mul(ia, ib).Add(id, ia).Div(id, ic)
        fmt.Printf("Big Int = %v\n", id)

        ja := big.NewRat(math.MaxInt64, 1994)
        jb := big.NewRat(-1994, math.MaxInt64)
        jc := big.NewRat(1, 10)
        jd := big.NewRat(1111, 2222)
        je := big.NewRat(1, 1)
        je.Mul(ja, jb).Add(je, jc).Mul(je, jd)
        fmt.Printf("Big Rat = %v\n", je)

}
