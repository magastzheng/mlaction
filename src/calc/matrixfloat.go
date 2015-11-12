package calc

import (
    //"fmt"
)

func GetSize(matrix [][]float64)(row, col int) {
    row = len(matrix)
    col = len(matrix[0])
    return
}

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

type Vector struct {
    len int
    data []float64
}

type Matrix interface {
    Dims() (r, c int)
    At(r, c int) float64
    T() Matrix
    Row(i int) Vector
    Column(i int) Vector
}

type FloatMatrix struct {
    rows int
    cols int
    data [][]float64
}

func (vec *Vector) Len() int {
    return vec.len
}

func (vec *Vector) Data() []float64 {
    return vec.data
}

func (vec *Vector) At(i int) float64 {
    if i >= vec.len || i < 0 {
        panic("The index is beyond the size of vector")
    }

    return vec.data[i]
}

func NewVector(a []float64) *Vector {
    return &Vector{
        len: len(a),
        data: a,
    }
}

func (fm *FloatMatrix) Data() [][]float64 {
    return fm.data
}

func (fm *FloatMatrix) Dims()(r, c int) {
    r = fm.rows
    c = fm.cols

    return
}

func (fm *FloatMatrix) At(r, c int) float64 {
    return fm.data[r][c]
}

func (fm *FloatMatrix) T() Matrix {
    m := new(FloatMatrix)
    m.data = make([][]float64, fm.cols)
    for i := 0; i < fm.cols; i++ {
        row := make([]float64, fm.rows)
        for j := 0; j < fm.rows; j++ {
            row[j] = fm.data[j][i]
        }
        m.data[i] = row
    }
    
    m.rows = fm.cols
    m.cols = fm.rows

    return m
}

func (fm *FloatMatrix) Row(i int) Vector {
    if i >= fm.rows || i < 0 {
        panic("The index is beyond the matrix row")
    }

    v := Vector{
        len: fm.cols,
        data: fm.data[i],
    }

    return v
}

func (fm *FloatMatrix) Column(i int) Vector {
    if i >= fm.cols || i < 0{
        panic("The index is beyond the matrix column")
    }

    v := Vector{
        len: fm.rows,
        data: make([]float64, fm.rows),
    }

    for j := 0; j < fm.rows; j++{
        v.data[j] = fm.data[j][i]
    }

    return v
}

func NewFloatMatrix(a [][]float64)(*FloatMatrix){
    return &FloatMatrix{
        rows: len(a),
        cols: len(a[0]),
        data: a,
    }
}

func AddMat(a, b Matrix) Matrix {
    ra, ca := a.Dims()
    rb, cb := b.Dims()
    if ra != rb || ca != cb {
        panic("Cannot add two matrix with different dims")
    }
    
    c := make([][]float64, ra)
    for row := 0; row < ra; row++ {
        temp := make([]float64, ca)
        for col := 0; col < ca; col++ {
            temp[col] = a.At(row,col)+b.At(row,col)
        }
        c[row] = temp
    }
    
    m := NewFloatMatrix(c)
    return m
}

func MultipyMat(a, b Matrix) Matrix {
    ra, ca := a.Dims()
    rb, cb := b.Dims()
    
    if ra == 0 || ca == 0 || rb == 0 || cb == 0{
        panic("Cannot multipy empty matrix")
    }

    if ca != rb {
        panic("Cannot multipy the two matrix for their size")
    }

    //initialize the result 2d array
    c := make([][]float64, ra)
    for i := 0; i < ra; i++{
        temp := make([]float64, cb)
        c[i] = temp
    }

    for i := 0; i < ra; i++ {
        for j := 0; j < cb; j++ {
            for k := 0; k < rb; k++ {
                c[i][j] += a.At(i,k) * b.At(k,j)
            }
        }
    }

    m := NewFloatMatrix(c)
    return m
}

func MultipyVec(a Matrix, b Vector) Vector {
    ra, ca := a.Dims()
    cb := b.Len()
    
    if ca != cb {
        panic("Cannot multipy the matrix and vector with different size")
    }

    c := make([]float64, ra)
    for i := 0; i < ra; i++ {
        for j := 0; j < ca; j++ {
            c[i] += a.At(i, j) * b.At(j)
        }
    }

    res := NewVector(c)

    return *res
}
