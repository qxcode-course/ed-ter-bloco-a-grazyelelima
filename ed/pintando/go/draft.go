package main

import (
	"fmt"
	"math"
)

func area_triangulo(a, b, c float64) float64 {
    p := (a + b + c) / 2.0
    return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func main() {
    var lado1, lado2, lado3 float64
    fmt.Scan(&lado1, &lado2, &lado3)
    result := area_triangulo(lado1, lado2, lado3)
    fmt.Printf("%.2f\n",result)
}
