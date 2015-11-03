package ml

import (
    "math"
    "fmt"
    //"reflect"
    "calc"
)

func LoadDataSet() ([][]string, []int) {
    postingList := [][]string{{"my", "dog", "has", "flea", "problems", "help", "please"},
                    {"maybe", "not", "take", "him", "to", "dog", "park", "stupid"},
                    {"my", "dalmation", "is", "so", "cute", "I", "love", "him"},
                    {"stop", "posting", "stupid", "worthless", "garbage"},
                    {"mr", "licks", "ate", "my", "steak", "how", "to", "stop", "him"},
                    {"quit", "buying", "worthless", "dog", "food", "stupid"},
                }
    classVec := []int{0, 1, 0, 1, 0, 1}

    return postingList, classVec
}

func isExisted(dataList []string, value string) bool {
    existed := false
    for _, d := range dataList {
        if d == value {
            existed = true
            break
        }
    }

    return existed
}

func CreateVocabList(dataSet [][]string)(vocabList []string) {
    for _, row := range dataSet {
        for _, data := range row {
            if !isExisted(vocabList, data) {
                vocabList = append(vocabList, data)
            }
        }
    }

    return
}

func lookupWord(vocabList []string, word string) (pos int) {
    pos = -1
    for i, d := range vocabList {
        if d == word {
            pos = i
            break
        }
    }

    return
}
func SetOfWords2Vec(vocabList []string, inputSet []string) (returnVec []int) {
    returnVec = make([]int, len(vocabList))
    for _,word := range inputSet {
        pos := lookupWord(vocabList, word)
        if pos >= 0 {
            returnVec[pos] = 1
        }else{
            fmt.Println("The word: ", word, " is not in my vocabulary!")
        }
    }

    return returnVec
}

func TrainNB0(trainMatrix [][]int, trainCategory []int)([]float64, []float64, float64) {
    numTrainDocs := len(trainMatrix)
    numWords := len(trainMatrix[0])
    sumCategory := calc.SumInt(trainCategory)
    
    pAbusive := float64(sumCategory) / float64(numTrainDocs)
    p0Num := make([]int, numWords)
    p1Num := make([]int, numWords)

    p0Denom := float64(0.0)
    p1Denom := float64(0.0)

    for i := 0; i < numTrainDocs; i++ {
        if trainCategory[i] == 1 {
            p1Num = calc.AddIntVec(p1Num, trainMatrix[i])
            p1Denom += float64(calc.SumInt(trainMatrix[i]))
        }else{
            p0Num = calc.AddIntVec(p0Num, trainMatrix[i])
            p0Denom += float64(calc.SumInt(trainMatrix[i]))        
        }
    }
    
    fmt.Println("p0Num: ", p0Num)
    fmt.Println("p1Num: ", p1Num)
    p1Vec := calc.DivIntVec(p1Num, p1Denom)
    p0Vec := calc.DivIntVec(p0Num, p0Denom)

    return p0Vec, p1Vec, pAbusive
} 

func LogVec(a []float64) (c []float64) {
    for _, v := range a {
        v = math.Log(v)
        c = append(c, v)
    }
    return
}

//Modified TrainNB0
//To reduce the impact of probability as 0
//use log()
func TrainNB1(trainMatrix [][]int, trainCategory []int)([]float64, []float64, float64) {
    numTrainDocs := len(trainMatrix)
    numWords := len(trainMatrix[0])
    sumCategory := calc.SumInt(trainCategory)
    
    pAbusive := float64(sumCategory) / float64(numTrainDocs)
    p0Num := make([]int, numWords)
    p1Num := make([]int, numWords)

    p0Denom := float64(2.0)
    p1Denom := float64(2.0)

    for i := 0; i < numTrainDocs; i++ {
        if trainCategory[i] == 1 {
            p1Num = calc.AddIntVec(p1Num, trainMatrix[i])
            p1Denom += float64(calc.SumInt(trainMatrix[i]))
        }else{
            p0Num = calc.AddIntVec(p0Num, trainMatrix[i])
            p0Denom += float64(calc.SumInt(trainMatrix[i]))        
        }
    }
    
    fmt.Println("p0Num: ", p0Num)
    fmt.Println("p1Num: ", p1Num)
    p1Vec := calc.DivIntVec(p1Num, p1Denom)
    p0Vec := calc.DivIntVec(p0Num, p0Denom)
    
    p0Vec = LogVec(p0Vec)
    p1Vec = LogVec(p1Vec)

    return p0Vec, p1Vec, pAbusive
} 

func ClassifyNB(vec2Classify []int, p0Vec []float64, p1Vec []float64, pClass1 float64) int {
    p1 := calc.MultipyIntVec(vec2Classify, p1Vec)
    p0 := calc.MultipyIntVec(vec2Classify, p0Vec)
    p1Val := calc.SumFloat(p1)
    p0Val := calc.SumFloat(p0)
    p1Val += math.Log(pClass1)
    p0Val += math.Log(1.0 - pClass1)
    
    if p1Val > p0Val {
        return 1
    }else{
        return 0
    }
}

func TestNB(){
    listOPosts, listClasses := LoadDataSet()
    myVocabList := CreateVocabList(listOPosts)
    
    trainMatrix := [][]int{}
    for _, row := range listOPosts {
        retVec := SetOfWords2Vec(myVocabList, row)
        trainMatrix = append(trainMatrix, retVec)
    }

    p0V, p1V, pAb := TrainNB0(trainMatrix, listClasses)
    testEntry := []string{"love", "my", "dalmation"}
    thisDoc := SetOfWords2Vec(myVocabList, testEntry)
    class := ClassifyNB(thisDoc, p0V, p1V, pAb)
    fmt.Println(testEntry, " is: ", class)

    testEntry = []string{"stupid", "garbage"}
    thisDoc = SetOfWords2Vec(myVocabList, testEntry)
    class = ClassifyNB(thisDoc, p0V, p1V, pAb)
    fmt.Println(testEntry, " is: ", class)
}
