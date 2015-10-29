package ml_test

import (
    "ml"
    "testing"
    "fmt"
)

func Test_knn(t *testing.T) {
    groups, labels := ml.CreateDataSet()
    inX := []float64{0.6, 2.0}
    result := ml.Classify(inX, groups, labels, 3)

    fmt.Println(result)
    inX = []float64{0, 0}
    result = ml.Classify(inX, groups, labels, 3)
    fmt.Println(result)
}
