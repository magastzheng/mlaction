package calc_test

import (
    "testing"
    "calc"
    "fmt"
    "math"
)

func Test_CalcRoot(t *testing.T){
    result := calc.CalcRoot(4, 2)
    fmt.Println(result)

    result = calc.CalcRoot(255, 5)
    fmt.Println(result)

    fmt.Println(math.Pow(3.02906, 5.0))
}


