package ml

import (
    //"fmt"
    "math"
)

func calcEntropy(prob float64) float64 {
    return -prob * math.Log2(prob)
}

//Input: the all the classes as an array
//Output: the entropy of the array
func CalcDataEntropy(labels []string) float64 {
    labelCountMap := make(map[string]int)
    for _, d := range labels {
        if v, ok := labelCountMap[d]; ok {
            labelCountMap[d] = v + 1
        }else{
            labelCountMap[d] = 1
        }
    }
    
    count := float64(len(labels))
    var entropy float64
    for _, v := range labelCountMap {
        prob := float64(v) / count
        temp := calcEntropy(prob)
        //fmt.Println(k, temp)
        entropy += temp
    }

    return entropy
}

func SplitDataSet(dataSet [][]int, axis int, value int) (retDataSet [][]int, rowIndex []int) {
    //retDataSet := [][]int{}

    for i, row := range dataSet {
        if axis >= len(row) {
            panic("Wrong axis is beyond the array space")
        }

        if row[axis] == value {
            retDataSet = append(retDataSet, row)
            rowIndex = append(rowIndex, i)
        }
    }

    return
}

func getFeatureList(dataSet [][]int, column int) (featList []int) {
    for i, totalRow := 0, len(dataSet) ; i < totalRow; i++ {
        data := dataSet[i][column]
        isExisted := false
        for _, v := range featList {
            if data == v {
                isExisted = true
                break
            }
        }

        if !isExisted{
            featList = append(featList, data)
        }
    }
    return
}

//Input the dataset of row*column d2 array, with its label with row size
//It will choose the best column(feature) as the split-standard. 
//According to the best information gain, which comes from baseEntropy-newEntropy
func ChooseBestFeatureToSplit(dataSet [][]int, labels []string) int {
    totalRow := len(dataSet)
    if totalRow != len(labels) {
        panic("The row of dataSet MUST be the same as size of labels")
    }
    
    numFeatures := len(dataSet[0])
    baseEntropy := CalcDataEntropy(labels)

    bestInfoGain := float64(0.0)
    bestFeature := -1
    for i := 0; i < numFeatures; i++ {
        featList := getFeatureList(dataSet, i)
        newEntropy := float64(0.0)
        
        for _, feature := range featList {
            subDataSet, rowIndex := SplitDataSet(dataSet, i, feature)
            prob := float64(len(subDataSet))/ float64(totalRow)
            subLabels := []string{}
            for _, row := range rowIndex {
                subLabels = append(subLabels, labels[row])
            }
            
            temp := CalcDataEntropy(subLabels)
            newEntropy += (prob * temp)
        }

        infoGain := baseEntropy - newEntropy
        
        if infoGain > bestInfoGain {
            bestInfoGain = infoGain
            bestFeature = i
        }
    }

    return bestFeature
}

//Selection algorithm
//Get the majority class in given input list as the target one
func MajorityCnt(classList []string) string {
    classCountMap := make(map[string]int)
    for _, class := range classList {
        if _, ok := classCountMap[class]; !ok {
            classCountMap[class] = 0
        }
        
        classCountMap[class] = classCountMap[class]+1
    }
    
    max := 0
    targetClass := ""
    for k, v := range classCountMap {
        if v > max {
            max = v
            targetClass = k
        }
    }

    return targetClass
}
