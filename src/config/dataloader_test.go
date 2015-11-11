package config_test

import (
    "testing"
    "config"
    "fmt"
)

func Test_LoadDataSet(t *testing.T){
    filename := "../resource/testSet.txt"
    dataSet, labels := config.LoadDataSet(filename)

    fmt.Println(len(dataSet))
    fmt.Println(len(labels))
}
