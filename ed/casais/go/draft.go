package main

import "fmt"

func main() {
    n := 0
    fmt.Scan(&n)

    animais := make([]int, n)
    for i := range animais {
        fmt.Scan(&animais[i])
    }

    solteiros := make(map[int]int)
    contCasais := 0

    for _, animal := range animais {
        if solteiros[-animal] != 0 {
            solteiros[-animal]--
            contCasais++
        } else {
            solteiros[animal]++
        }
    }

    fmt.Println(contCasais)
}
