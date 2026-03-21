package main
import "fmt"

type Gomo struct {
    a int 
    b int
}

func main() {

    var q int 
    var direcao string

    fmt.Scan(&q)
    fmt.Scan(&direcao)

    cobra := make([]Gomo, q)

    for i := 0; i < q; i++ {
        fmt.Scan(&cobra[i].a, &cobra[i].b)
    }

    posicaoAnteriores := make([]Gomo, q)
    copy(posicaoAnteriores, cobra)

    switch direcao {
        case "L":
            cobra[0].a--
        case "R":
            cobra[0].a++
        case "U":
            cobra[0].b--
        case "D":
            cobra[0].b++        
    }

    for i := 1; i < q; i++{
        cobra[i] = posicaoAnteriores[i-1]
    }

    for _, g := range cobra {
        fmt.Printf("%d %d\n", g.a, g.b)
    }


}
