package ml

import (
    "fmt"
    "math"
)

func Sigmoid(x float64) float64 {
    x = math.Exp(0-x)
    fmt.Println(x)
    x = float64(1.0)+x
    x = 1.0 / x
    return x
}

func LoadDataSet_log()([][]float64, []int) {
    dataMat := [][]float64{}
    labelMat := []int{}

    dataMat = append(dataMat, []float64{1.0, 1.5, 2.6})
    dataMat = append(dataMat, []float64{1.0, 2.5, 1.6})
    dataMat = append(dataMat, []float64{1.0, 3.5, 0.6})
    dataMat = append(dataMat, []float64{1.0, 4.5, 5.6})
    dataMat = append(dataMat, []float64{1.0, 5.5, 7.6})

    labelMat = append(labelMat, 1)
    labelMat = append(labelMat, 0)
    labelMat = append(labelMat, 1)
    labelMat = append(labelMat, 0)
    labelMat = append(labelMat, 1)

    return dataMat, labelMat
}


//
//func GradAscent(dataMatIn [][]float64, classLabels []int) (weights []float64){
//    alpha := 0.001
//    maxCycles := 500
//
//}
