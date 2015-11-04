package calc

import(
    "fmt"
    "math"
    "math/rand"
)

//Use Newton iteration method to calculate root of x**n = a
//x1=(1-1/n)x0+a*x0/(x0**(n-1))
func CalcRoot(a, n int) float64 {
    e := float64(0.0002)
    seed := int64(a / n)
    r := rand.New(rand.NewSource(seed))
    x0 := r.Float64()
    x1 := float64(0.0)
    result := x0 - x1
    fmt.Println("x0,x1,result: ", x0, x1, result)
    for {
        if result > (0-e) && result < e {
            break;
        }
        
        temp1 := float64(1.0) / float64(n)
        temp1 = float64(1.0) - temp1
        temp1 *= x0
         
        n1pow := math.Pow(x0, float64(n - 1))
        n1pow *= float64(n)
        temp2 := float64(a) / n1pow

        x1 = temp1 + temp2
        result = x0 - x1

        x0 = x1
    }

    return x1
}
