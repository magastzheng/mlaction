package calc

import (
    //"fmt"
)

func AddMatrixInt(a [][]int, b [][]int) (c [][]int) {
    if len(a) != len(b){
        panic("Wrong size of matrix")
    }
    
    for i, row := 0, len(a); i < row; i++ {
        ai := a[i]
        temp := make([]int, len(ai))
        for j, col := 0, len(ai); j < col; j++{
            temp[j] = ai[j]+b[i][j]
        }
        c = append(c, temp)
    }

    return
}

func MultipyMatrixInt(a [][]int, b [][]int)(c [][]int) {
    arow, brow := len(a), len(b)
    if arow == 0 || brow == 0 {
        panic("Cannot multipy empty matrix")
    }

    acol, bcol := len(a[0]), len(b[0])
    if acol != brow{
        panic("Two wrong size matrix to multipy")
    }
    
    c = make([][]int, arow)
    for i := 0; i < arow; i++ {
        temp := make([]int, bcol)
        c[i] = temp
    }
    
    //a[M][K]*b[K][N] => c[M][N] 
    //a[arow][acol]*b[acol][bcol] => c[arow][bcol]
    for i := 0; i < arow; i++ {
        for j := 0; j < bcol; j++ {
            for k := 0; k < brow; k++ {
                c[i][j] += a[i][k] * b[k][j]
            }
        }
    }

    return
}

func TransposeMatrixInt(a [][]int)(c [][]int){
    row := len(a)
    if row == 0 {
        panic("Cannot transpose empty matrix")
    }
    col := len(a[0])
    c = make([][]int, col)
    for i := 0; i < col; i++ {
        rows := make([]int, row)
        for j := 0; j < row; j++ {
            rows[j] = a[j][i]
        }
        c[i] = rows
    }

    return
}

