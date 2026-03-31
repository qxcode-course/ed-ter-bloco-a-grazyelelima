package main
import "fmt"

func saidaVivos(vivos []int, donoDaEspada int) {
    fmt.Print("[ ")
    for _, p := range vivos {
        if p == donoDaEspada {
            fmt.Printf("%d> ", p)
        } else {
            fmt.Printf("%d ", p)
        }
    }
    fmt.Println("]")
}

func main() {

    var n, e int
    fmt.Scan(&n, &e)

    vivos := make([]int, n)

    for i := 0; i < n; i++ {
        vivos[i] = i + 1
    }

    pos := 0

    for i, v := range vivos {
        if v == e {
            pos = i
            break
        }
    }

    for len(vivos) > 1 {
        saidaVivos(vivos, vivos[pos])

        morre := (pos + 1) % len(vivos)

        vivos = append(vivos[:morre], vivos[morre+1:]...)

        if morre < pos {
            pos--
        }

        pos = (pos + 1) % len(vivos)
    }

    saidaVivos(vivos, vivos[0])

    

}
