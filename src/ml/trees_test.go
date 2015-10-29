package ml_test

import (
    "ml"
    "testing"
    "fmt"
)

func Test_CalcDataEntropy(t *testing.T) {
    groups := []string{"Yes", "Yes", "No", "Yes", "No", "No", "Yes", "No", "Yes", "Maybe"}
    entropy := ml.CalcDataEntropy(groups)
    fmt.Println(entropy)
}
 
func Test_SplitDataSet(t *testing.T) {
    dataSet := [][]int{{1, 0, 2}, {1, 1, 0}, {0, 2, 1}, {0, 1, 2}, {2, 1, 0}}
    fmt.Println("source: ", dataSet)
    axis := 0
    value := 1
    result := ml.SplitDataSet(dataSet, axis, value)
    fmt.Println("Index: ", axis, " Value: ", value)
    fmt.Println( result )

    axis = 1
    result =ml.SplitDataSet(dataSet, axis, value)
    fmt.Println("Index: ", axis, " Value: ", value)
    fmt.Println( result )

    axis = 1
    value = 2
    result =ml.SplitDataSet(dataSet, axis, value)
    fmt.Println("Index: ", axis, " Value: ", value)
    fmt.Println( result )
}
