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
