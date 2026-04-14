package main

import "fmt"

func main() {

    var n int
    fmt.Scan(&n)

    animais := make([]int , n)

    for i := 0; i < n; i++ {
        fmt.Scan(&animais[i])
    }

    solteiros := make(map[int]int)

    numCasais := 0
    for _, animal := range animais {
        if solteiros[-animal] > 0 {
            solteiros[-animal]--
            numCasais++
        } else {
            solteiros[animal]++
        }
    }

    fmt.Println(numCasais)

}
