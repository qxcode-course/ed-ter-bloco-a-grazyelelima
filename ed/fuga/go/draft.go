package main
import "fmt"
func main() {
    var h, p, f, d int

    fmt.Scan(&h, &p, &f, &d)

    var vaiFugir int = 1

    for f != h && vaiFugir == 1 {
        f = (f + d + 16) % 16
        if f == p {
            vaiFugir = 0
        }

    }

    if vaiFugir == 1 {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}
