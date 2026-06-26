package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func backtrack(letras []string, visitados []bool, atual []string) {
    if len(atual) == len(letras) {
        fmt.Println(strings.Join(atual, ""))
        return
    }

    for i := 0; i < len(letras); i++ {
        if !visitados[i] {
            visitados[i] = true
            atual = append(atual, letras[i])

            backtrack(letras, visitados, atual)

            atual = atual[:len(atual)-1]
            visitados[i] = false
        }
    }
}


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    if !scanner.Scan() {
        return 
    }

    entrada := strings.TrimSpace(scanner.Text())
    if entrada == "" {
        return 
    }

    letras := strings.Split(entrada, "")
    sort.Strings(letras)

    visitados := make([]bool, len(letras))
    atual := make([]string, 0)

    backtrack(letras, visitados, atual)
}