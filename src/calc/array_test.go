package calc_test

import (
    "testing"
    "calc"
    "fmt"
    //"reflect"
)

func Test_Sum(t *testing.T) {
    a := []int{1, 2, 3}
    result := calc.Sum(a)
    //var b interface{}
    //b = a
    //fmt.Println(reflect.TypeOf(b))
    //fmt.Println(b)
    fmt.Println(result)
    b := int(result.(int64))
    fmt.Println(b)
}

func Test_Sum_Float(t *testing.T) {
    a := []float32{1.5, 2.3, 3.6}
    result := calc.Sum(a)
    //var b interface{}
    //b = a
    //fmt.Println(reflect.TypeOf(b))
    //fmt.Println(b)
    fmt.Println(result)
    fmt.Println(result.(float64))
    b := float32(result.(float64))
    fmt.Println(b)

    s := float32(0.0)
    for _, v := range a {
        s += v
    }
    fmt.Println(s)
}

func Test_SumInt(t *testing.T) {
    a := []int{1, 2, 3}
    result := calc.SumInt(a)
    if result != 6 {
        t.Fatalf("Sum with the wrong value")
    }
}

func Test_SumFloat(t *testing.T) {
    a := []float64{1.23, 3.24, 4.65}
    result := calc.SumFloat(a)
    t.Log(result)
    if result > 9.2 || result < 9.0 {
        t.Fatalf("Sum with the wrong value")
    }
}

func Test_AddIntVec(t *testing.T) {
    a := []int{1, 2, 3}
    b := []int{4, 5, 6}
    result := calc.AddIntVec(a, b)
    if len(result) != len(a) {
        t.Fatalf("Add wrong")
    }

    for i, count := 0, len(a); i < count; i++ {
        if result[i] != a[i]+b[i] {
            t.Fatalf("Add to wrong result")
        }
    }
}

func Test_DivIntVec(t *testing.T) {
    a := []int{4, 7, 9}
    d := float64(2.0)
    result := calc.DivIntVec(a, d)
    
    if len(result) != len(a) {
        t.Fatalf("Div to the wrong size")
    }
}

func Test_MultipyIntVec(t *testing.T) {
    a := []int{1, 2, 3}
    b := []float64{2.25, 3.14, 4.67}
    result := calc.MultipyIntVec(a, b)

    if len(result) != len(a) {
        t.Fatalf("Multipy to get wrong size")
    }
}
