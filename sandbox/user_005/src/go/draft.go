package main
import "fmt"

func fatorial(value int) int {
    if value == 1 {
        return 1
    }
    fmt.Println(value)
    coletado := value * fatorial(value-1)
    fmt.Println(coletado)
    return coletado
}
func main() {
    num := 0
    fmt.Scan(&num)
    fatorial(num)
}