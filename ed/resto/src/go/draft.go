package main

import "fmt"

func divisao(x int) {
    if x == 0 {
        return 
    }

    divisao(x/2)

    fmt.Printf("%d %d\n", x/2, x%2)
}

func main() {
    
    n := 0
    fmt.Scan(&n)

    divisao(n)
    
}
