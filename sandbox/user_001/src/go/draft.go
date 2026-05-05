package main

import "fmt"

func main() {

    var nome string = "Grazy"
    idade := 21
    altura := 1.67
    gostaDeCodar := true

    fmt.Printf("%s tem %d anos, %.2fm e gosta de codar em go: %v \n", nome, idade, altura, gostaDeCodar)
}
