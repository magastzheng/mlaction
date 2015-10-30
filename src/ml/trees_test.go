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

    groups = []string{"Yes", "Yes"}
    entropy = ml.CalcDataEntropy(groups)
    fmt.Println(entropy)

    groups = []string{"Yes", "No"}
    entropy = ml.CalcDataEntropy(groups)
    fmt.Println(entropy)
}
 
func Test_SplitDataSet(t *testing.T) {
    dataSet := [][]int{{1, 0, 2}, {1, 1, 0}, {0, 2, 1}, {0, 1, 2}, {2, 1, 0}}
    fmt.Println("source: ", dataSet)
    axis := 0
    value := 1
    result, rowIndex := ml.SplitDataSet(dataSet, axis, value)
    fmt.Println("Index: ", axis, " Value: ", value)
    fmt.Println( result, rowIndex )

    axis = 1
    result, rowIndex =ml.SplitDataSet(dataSet, axis, value)
    fmt.Println("Index: ", axis, " Value: ", value)
    fmt.Println( result, rowIndex )

    axis = 1
    value = 2
    result, rowIndex =ml.SplitDataSet(dataSet, axis, value)
    fmt.Println("Index: ", axis, " Value: ", value)
    fmt.Println( result , rowIndex )
}

func Test_ChooseBestFeatureToSplit(t *testing.T) {
    dataSet := [][]int{{1, 0, 2}, {1, 1, 0}, {0, 2, 1}, {0, 1, 2}, {2, 1, 0},{2, 1, 3}, {0, 3, 1}}
    labels := []string{"Yes", "Yes", "No", "Yes", "No", "Maybe", "No"}
    
    fmt.Println(dataSet)
    fmt.Println(labels)
    bestFeature := ml.ChooseBestFeatureToSplit(dataSet, labels)
    fmt.Println("testtest")
    fmt.Println(bestFeature)
}

func Test_MajorityCnt(t *testing.T) {
    groups := []string{"A", "B", "C", "A", "B", "A"}
    result := ml.MajorityCnt(groups)

    fmt.Println(result)
}
