package calc_test

import (
    "testing"
    "calc"
    "fmt"
)

func Test_MatrixFloat(t *testing.T){
    a := [][]float64{{1.2, 1.3}, {2.3, 4.5}}
    m := calc.NewFloatMatrix(a)
    fmt.Println(m)

    r, c := m.Dims()
    if r != len(a) || c != len(a[0]) {
        t.Fatalf("Calculate wrong dims")
    }
    
    el := m.At(1, 1)
    fmt.Println(el)
    if el < 4.4999999 || el > 4.50000001 {
        t.Fatalf("Get the wrong value")
    }

    mt := m.T()
    fmt.Println(mt) 
}

func Test_MatrixFloat_RowCol(t *testing.T){
    a := [][]float64{{1.2, 1.3}, {2.3, 4.5}, {3.5, 7.8}}
    m := calc.NewFloatMatrix(a)
    fmt.Println(m)

    r, c := m.Dims()
    if r != len(a) || c != len(a[0]) {
        t.Fatalf("Calculate wrong dims")
    }
    
    row1 := m.Row(1)
    fmt.Println(row1)

    col1 := m.Column(1)
    fmt.Println(col1)
}


func Test_MatrixFloat_Add(t *testing.T){
    a := [][]float64{{1.2, 1.3}, {2.3, 4.5}, {3.5, 7.8}}
    b := [][]float64{{-1.1, -2.3}, {2.3, 5.5}, {5.3, 8.2}}
    m1 := calc.NewFloatMatrix(a)
    m2 := calc.NewFloatMatrix(b)
    fmt.Println(m1, m2)
    
    m3 := calc.AddMat(m1, m2)
    fmt.Println(m3)
}

func Test_MatrixFloat_MultipyMat(t *testing.T){
    a := [][]float64{{1.2, 1.3}, {2.3, 4.5}, {3.5, 7.8}}
    b := [][]float64{{2.1, 1.2, 3.2}, {2.3, 5.5, 5.3}}
    m1 := calc.NewFloatMatrix(a)
    m2 := calc.NewFloatMatrix(b)
    fmt.Println(m1, m2)
    
    m3 := calc.MultipyMat(m1, m2)
    fmt.Println(m3)
}

func Test_MatrixFloat_MultipyVec(t *testing.T){
    a := [][]float64{{1.2, 1.3}, {2.3, 4.5}, {3.5, 7.8}}
    b := []float64{2.1, 1.2}
    m1 := calc.NewFloatMatrix(a)
    v1 := calc.NewVector(b)
    fmt.Println(m1, v1)
    
    v2 := calc.MultipyVec(m1, *v1)
    fmt.Println(v2)
}
