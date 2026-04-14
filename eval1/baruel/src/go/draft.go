package main
import "fmt"
func main() {
    
    var qtd_album, qtd_possui int
    fmt.Scan(&qtd_album, &qtd_possui)

    unicas := make(map[int]bool)
    figuras := make([]int, qtd_possui)

    for i := range figuras {
        fmt.Scan(&figuras[i])
    }

    repetidas := make([]int, 0, qtd_possui)
    faltantes := make([]int, 0)

    for _, figura := range figuras {
        if unicas[figura] {
            repetidas = append(repetidas, figura)
        } else {
            unicas[figura] = true
        }
    }

    for i := 1; i <= qtd_album; i++ {
        if !unicas[i] {
            faltantes  = append(faltantes, i)
        }
    }

    saida := fmt.Sprintf("%v", repetidas)
    if saida == "[]" {
        fmt.Println("N")
    } else {
        fmt.Println(saida[1 : len(saida) -1])
    }

    saida = fmt.Sprintf("%v", faltantes)
    if saida == "[]" {
        fmt.Println("N")
    } else {
        fmt.Println(saida[1 : len(saida) -1])
    }
    
    //Lembrei como formata a saida KKKKK

}
