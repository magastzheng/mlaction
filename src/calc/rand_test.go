package calc_test

import (
    "testing"
    "calc"
    "fmt"
)

func Test_RandInt(t *testing.T) {
    //for i := 0; i < 1000; i++ {
        n := calc.RandInt(10000)
        //fmt.Println(n)
    //}

    if n < 0 || n > 10000 {
        t.Fatalf("The random value is beyond the range")
    }
}

func Test_RandIntArray(t *testing.T) {
    a := calc.RandIntArray(10)
    
    b := calc.RandIntArray(10)
    if len(a) != len(b) {
        t.Fatalf("The two random arrays have different size")
    }else{
        for i := 0; i < len(a); i++ {
            if a[i] != b[i] {
                t.Fatalf("The two random arrays have different sequence")
            }
        }
    }
}

func Test_SelectJrand(t *testing.T) {
    result := calc.SelectJrand(3, 20)
    fmt.Println(result)
    if result == 3 || result >= 20 {
        t.Fatalf("The value is wrong")
    }
}

func Test_ClipAlpha(t *testing.T) {
    aj := 12.5
    H := 11.3
    L := 2.5
    result := calc.ClipAlpha(aj, H, L)
    fmt.Println(result)
    if result > H || result < L {
        t.Fatalf("The value is beyond the given range")
    }
}
