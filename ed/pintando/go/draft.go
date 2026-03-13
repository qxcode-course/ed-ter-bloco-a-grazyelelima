package main

import (
    "fmt"
    "math")

func main() {
    var a float64
    var b float64 
    var c float64

    var area float64

    fmt.Scan(&a, &b, &c)

    var p float64 = float64(a + b + c) / 2.0

    area = math.Sqrt(p * (p - float64(a)) * (p - float64(b)) * (p - float64(c)))

    fmt.Printf("%.2f\n", area)
}
