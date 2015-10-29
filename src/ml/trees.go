package ml

import (
    "fmt"
    "math"
)

func calcEntropy(prob float64) float64 {
    return -prob * math.Log2(prob)
}

//Input: the all the classes as an array
//Output: the entropy of the array
func CalcDataEntropy(data []string) float64 {
    dataCountMap := make(map[string]int)
    for _, d := range data {
        if v, ok := dataCountMap[d]; ok {
            dataCountMap[d] = v + 1
        }else{
            dataCountMap[d] = 1
        }
    }
    
    count := float64(len(data))
    var entropy float64
    for k, v := range dataCountMap {
        prob := float64(v) / count
        temp := calcEntropy(prob)
        fmt.Println(k, temp)
        entropy += temp
    }

    return entropy
}

func SplitDataSet(dataSet [][]int, axis int, value int) [][]int {
    retDataSet := [][]int{}

    for _, row := range dataSet {
        if axis >= len(row) {
            panic("Wrong axis is beyond the array space")
        }

        if row[axis] == value {
            retDataSet = append(retDataSet, row)
        }
    }

    return retDataSet
}
