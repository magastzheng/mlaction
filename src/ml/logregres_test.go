package ml_test

import(
    "testing"
    "ml"
    "fmt"
)

func Test_Sigmoid(t *testing.T){
    result := ml.Sigmoid(float64(2))
    fmt.Println(result)
}
