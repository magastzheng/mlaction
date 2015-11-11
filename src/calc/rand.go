package calc

import (
    "math/rand"
    "time"
    //"fmt"
)

func RandInt(n int) int {
   return rand.Intn(n)
}

func RandIntArray(n int) []int {
    seed := time.Now().UnixNano()
    //fmt.Println(seed)
    r := rand.New(rand.NewSource(seed))

    a := make([]int, 0)
    for i := 0; i < n; i++ {
        a = append(a, r.Intn(n))
    }

    return a
}

//Select a value j which will be in [0, m) but is not equal to i
func SelectJrand(i, m int) int {
    j := i
    for {
        j = rand.Intn(m)
        if j != i {
            break
        }
    }

    return j
}

//Adjust the value of aj. It will ensure the aj in arange: [L, H]
func ClipAlpha(aj, H, L float64) float64 {
    if aj > H {
        aj = H
    }
    if L > aj {
        aj = L
    }

    return aj
}
