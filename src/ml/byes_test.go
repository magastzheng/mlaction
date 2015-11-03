package ml_test

import (
    "fmt"
    "ml"
    "testing"
)

func Test_LoadDataSet(t *testing.T) {
    dataSet, classVec := ml.LoadDataSet()
    fmt.Println(len(dataSet), dataSet)
    fmt.Println(len(classVec), classVec)
}

func Test_CreateVocabList(t *testing.T){
    dataSet, _ := ml.LoadDataSet()
    vocabList := ml.CreateVocabList(dataSet)

    fmt.Println(len(vocabList), vocabList)
}

func Test_SetOfWords2Vec(t *testing.T){
    dataSet, _ := ml.LoadDataSet()
    vocabList := ml.CreateVocabList(dataSet)
    inputVec := []string{"my", "dog", "is", "not", "like", "that"}

    retVec := ml.SetOfWords2Vec(vocabList, inputVec)
    fmt.Println(retVec)
}

func Test_TrainNB0(t *testing.T) {
    dataSet, categories := ml.LoadDataSet()
    vocabList := ml.CreateVocabList(dataSet)
    trainMatrix := [][]int{}
    for i, row := range dataSet {
        retVec := ml.SetOfWords2Vec(vocabList, row)
        fmt.Println(retVec , "======", i)
        //trainMatrix[i] = make([]int, len(retVec))
        //trainMatrix[i] = append(trainMatrix[i], retVec...)
        trainMatrix = append(trainMatrix, retVec)
    }
    fmt.Println(trainMatrix)
    fmt.Println(categories)
    p0Vec, p1Vec, pAbusive := ml.TrainNB0(trainMatrix, categories)
    fmt.Println(vocabList)
    fmt.Println("p0: ", p0Vec)
    fmt.Println("p1: ", p1Vec)
    fmt.Println("pAbusive: ",pAbusive)
}

func Test_TetstNB(t *testing.T) {
    ml.TestNB()
}
