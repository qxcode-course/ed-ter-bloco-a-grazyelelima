package main
import (
    "fmt"
    "math"
)
func main() {
    var n int 

    fmt.Scan(&n)

    var ganhador int = -1
    var melhorDif int = 101

    for i := 0; i < n; i++ {
        var a, b int 
        fmt.Scan(&a, &b)

        dif := int(math.Abs(float64(a - b)))

        

        if a < 10 || b < 10 {
            continue
        } 

        if dif < melhorDif {
            melhorDif = dif 
            ganhador = i
        }

    }

    if ganhador == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(ganhador)
    }

}
