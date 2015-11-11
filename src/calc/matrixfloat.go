package calc

import (
    //"fmt"
)


//Calculate the |A| expand the first row
//determinant
func getA(arcs [][]float64, n int)float64{
    row := len(arcs)
    col := len(arcs)

    if row != col {
        panic("Cannot calcute the matrix with different row an column.")
    }

    if n == 1 {
        return arcs[0][0]
    }

    ans := float64(0.0)

    var i, j, k int
    for i = 0; i < n; i++ {
        tempmatrix := make([][]float64, n - 1)
        for j = 0; j < n - 1; j++ {
            temprow := make([]float64, n - 1)
            for k = 0; k < n-1; k++ {
                index := k
                if k >= i {
                    index = k + 1
                }
                temprow[k] = arcs[j+1][index]
            }
            tempmatrix[j] = temprow
        }

        determinant := getA(tempmatrix, n - 1)

        if i % 2 == 0 {
            ans += arcs[0][i] * determinant
        } else {
            ans -= arcs[0][i] * determinant
        }
    }
    
    return ans
}

//adjoint A*
//func getAdjoint(arcs [][]float64) [][]float64 {
//    
//}
