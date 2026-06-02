package main

import (
	"fmt"
	"strconv"
    "strings"
)


func imprimindoVetor(vet []int) {
    primos := make([]string, len(vet))
    for i, v := range vet {
        primos[i] = strconv.Itoa(v)
    }

    v := strings.Join(primos, ", ")
    fmt.Printf("[%s]\n", v)
}

func priminhos(n int, atual int, lista []int) []int {
    if len(lista) == n {
        return lista
    }

    if eh_primo(atual, 2) {
        lista = append(lista, atual)
    }

    return priminhos(n, atual+1, lista)
}

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	if x < 2 {
		return false
	}

	if x == 2 {
		return true
	}

	if div * div > x {
		return true
	}

	if x % div == 0 {
		return false
	}

	return eh_primo(x, div + 1)
}

func main() {
	var x int
	fmt.Scan(&x)
	listaInicial := make([]int, 0)
    result := priminhos(x, 2, listaInicial)
    imprimindoVetor(result)
}
