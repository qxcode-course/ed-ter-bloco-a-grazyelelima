package main

import "fmt"

func verificarSeFugiu(f, d, h, p int) bool {
    fugiuOuNao := false
    for {
        f += d

        if f <= 0 {
            f = 15
        }

        if f > 15 {
            f = 0
        }

        if f == h {
            fugiuOuNao = true
            break
        }

        if f == p {
            break
        }
    }
    return fugiuOuNao
}

func main() {
    var h, p, f, direcao int
    fmt.Scan(&h, &p, &f, &direcao)

    if verificarSeFugiu(f, direcao, h, p) {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}    

