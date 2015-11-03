package calc_test

import (
    "testing"
    "calc"
    "fmt"
)

func Test_Add(t *testing.T) {
    a, _ := calc.Add(10, 20)
    fmt.Println(a)

    var b int64
    b = a.(int64)
    fmt.Println(b)
}
