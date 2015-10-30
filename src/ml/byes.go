package ml

import (
    "fmt"
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
