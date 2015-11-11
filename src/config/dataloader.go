package config

import (
    "fmt"
    "os"
    //"io"
    "bufio"
    "strings"
    "strconv"
)

func LoadDataSet(filename string) ([][]float64, []int) {
    file, err := os.Open(filename)
    if err != nil {
        panic("Cannot open file: " + filename)
    }
    defer file.Close()
    
    dataSet := make([][]float64, 0)
    labels := make([]int, 0)
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        column := strings.Split(line, "\t")
        
        dataRow := make([]float64, 0)
        for i, count := 0, len(column); i < count; i++ {
            v := column[i]
            if i == count - 1 {
                if label, err := strconv.ParseInt(v, 10, 32); err == nil {
                    labels = append(labels, int(label))
                } else {
                    panic(err)
                }
            } else {
                if value, err := strconv.ParseFloat(v, 64); err == nil {
                    dataRow = append(dataRow, value)
                }else{
                    fmt.Println("Cannot convert float64: ", v)
                    panic(err)
                }
            }
        }

        dataSet = append(dataSet, dataRow)
    }

    return dataSet, labels
}

