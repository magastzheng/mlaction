package calc

import (
    //"math"
    //"fmt"
    "reflect"
)

//Sum of the slice
func Sum(a interface{}) interface{} {
    var sum interface{}
    switch reflect.TypeOf(a).Kind() {
        case reflect.Slice:
            s := reflect.ValueOf(a)
            sum = s.Index(0).Interface()
            for i := 1; i < s.Len(); i++ {
                //fmt.Println(s.Index(i), s.Index(i).Interface())
                sum, _ = Add(sum, s.Index(i).Interface())
            }
        default:
    }

    return sum
}

//func Sum(a []interface{}) interface{} {
func sum(a []interface{}) interface{} {
    var sum interface{} 
    for _, v := range a {
        sum, _ = Add(sum, v)
    }

    return sum
}

func SumInt(a []int) int {
    c := Sum(a)
    return int(c.(int64))
}

func SumFloat(a []float64) float64 {
    c := Sum(a)
    return c.(float64)
}

func AddIntVec(a []int, b []int) (c []int){
    if len(a) != len(b) {
        panic("The two vectors with different length")
    }
    
    for i, count := 0, len(a); i < count; i++ {
        c = append(c, a[i]+b[i])
    }

    return c
}

func DivIntVec(a []int, denom float64) (c []float64) {
    for i, count := 0, len(a); i < count; i++ {
        c = append(c, float64(a[i]) / denom)
    }

    return c
}

func MultipyIntVec(a []int, b []float64) (c []float64) {
    if len(a) != len(b) {
        panic("The two vectors with different size")
    }

    for i, count := 0, len(a); i < count; i++ {
        temp := float64(a[i]) * b[i]
        c = append(c, temp)
    }

    return
}
