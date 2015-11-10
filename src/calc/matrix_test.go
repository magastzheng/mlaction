package calc_test

import (
    "calc"
    "testing"
    "fmt"
)

func Test_AddMatrixInt(t *testing.T) {
    a := [][]int{{1, 2, 3}, {4, 5, 6}}
    b := [][]int{{1, 2, 3}, {4, 5, 6}}
    result := calc.AddMatrixInt(a, b)
    fmt.Println(result)
}

func Test_MultipyMatrixInt(t *testing.T) {
    a := [][]int{{1, 2, 3}, {4, 5, 6}}
    b := [][]int{{3, 2, 1}, {2, 2, 1}, {3, 1, 2}}

    result := calc.MultipyMatrixInt(a, b)
    fmt.Println(result)
}

func Test_TransposeMatrixInt(t *testing.T) {
    a := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}}
    result := calc.TransposeMatrixInt(a)
    fmt.Println(result)
}
