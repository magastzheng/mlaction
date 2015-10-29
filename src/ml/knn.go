package ml

import (
    "fmt"
    "math"
    "sort"
)

type Distance struct {
    Index int
    Value float64
}

type DistanceWrapper struct {
    d []Distance
    by func(p, q *Distance) bool
}

func (dw DistanceWrapper) Len() int {
    return len(dw.d)
}

func (dw DistanceWrapper) Swap(i, j int) {
    dw.d[i], dw.d[j] = dw.d[j], dw.d[i]
}

func (dw DistanceWrapper) Less(i, j int) bool {
    return dw.by(&dw.d[i], &dw.d[j])
}

func CreateDataSet() ([][]float64, []string) {
    group := [][]float64{}
    g1 := []float64{1.0, 1.1}
    g2 := []float64{1.0, 1.0}
    g3 := []float64{0, 0}
    g4 := []float64{0, 0.1}
    group = append(group, g1)
    group = append(group, g2)
    group = append(group, g3)
    group = append(group, g4)
    labels := []string{"A", "A", "B", "B"}
    
    fmt.Println(group)
    return group, labels
}

func Classify(inX []float64, dataSet [][]float64, labels []string, k int) string {
    dataSetSize := len(dataSet)
    inXSize := len(inX)

    distances := make([]Distance, 0)
    //Loop to calculate the distance
    for i := 0; i < dataSetSize; i++{
        caseData := dataSet[i]
        caseDataLen := len(caseData)
        if inXSize != caseDataLen {
            panic("Input size is not equal to dataSet size")
        }
        
        var d float64
        for j, value := range caseData {
            temp := (inX[j] - value) * (inX[j] - value)
            d += temp
        }
        fmt.Print("d=", d)
        d = math.Sqrt(d)
        fmt.Println("sqrt=", d)

        distance := Distance{i, d}
        distances = append(distances, distance)
    }

    sort.Sort(DistanceWrapper{distances, func (p, q *Distance) bool {
        return p.Value < q.Value
    }})
    fmt.Println(distances)
    classCountMap := make(map[string]int)
    
    //statistics the class count
    if k > len(labels) {
        panic("K should be less than length of labels")
    }

    for i := 0; i < k; i++ {
        pos := distances[i].Index
        label := labels[pos]
        if v, ok := classCountMap[label]; ok {
            classCountMap[label] = v+1
        }else{
            classCountMap[label] = 1
        }
    }

    //get maximum class
    var key string
    var total int
    for k, v := range classCountMap {
        if v > total {
            key = k
            total = v
        }
    }

    return key
}


