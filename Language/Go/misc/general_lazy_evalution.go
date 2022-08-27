package main

import (
        "fmt"
)

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func main() {
        oddFunc := func(state Any) (Any, Any) {
                //f(n) = f(n-1) + f(n-2)
                os := state.(int) + 1
                ns := os + 1
                return os, ns  
        }
        odd := BuildLazyIntEvaluator(oddFunc, 0)

        for i := 0; i < 10; i++ {
                fmt.Printf("%vth fib number is: %v\n", i, odd())
        }
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
        retValChan := make(chan Any)
        loopFunc := func() {
                var actState Any = initState
                var retVal Any
                for {
                        retVal, actState = evalFunc(actState)
                        retValChan <- retVal 
                }
        }
        retFunc := func() Any {
                return <- retValChan 
        }
        go loopFunc()
        return retFunc 
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
        ef := BuildLazyEvaluator(evalFunc, initState)
        return func() int {
                return ef().(int)
        }
}


