package main

import (
	"fmt"
	"math"
        )

func area(a, b, c float64) float64 {
    p := (a + b + c) / 2.0
    return math.Sqrt(p  * (p - a) * (p - b) * (p - c))
}

func main() {
    
    var l1, l2, l3 float64
    fmt.Scan(&l1, &l2, &l3)

    result := area(l1, l2, l3)
    fmt.Printf("%.2f\n", result)

}
