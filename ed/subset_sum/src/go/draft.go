package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func backtrack(idx int, somaAtual int, alvo int, n int, arr []int) bool {

    //caso base 1: se a somaAtual, ou seja, acumulado pelas escolhas anteiores for igual ao alvo, significa que o subconjubto atual funciona
    if somaAtual == alvo {
        return true
    }

    //caso base 2: se a soma já passou do alvo, qualquer numero adicionado dps vai aumentar a soma ainda mais, logo é impossível chegar ao alvo se já passou dele, voltameos. Se o indice chegou ao fim e não parou no caso um, a soma não bateu!
    if somaAtual > alvo || idx >= n {
        return false
    }

    if backtrack(idx+1, somaAtual+arr[idx], alvo, n, arr) {
        return true
    }

    if backtrack(idx+1, somaAtual, alvo, n, arr) {
        return true
    }

    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var l1 []string
    for scanner.Scan() {
        l1 = strings.Fields(scanner.Text())
        if len(l1) > 0 {
            break
        }
    }

    if len(l1) < 2 {
        return 
    }

    n, _ := strconv.Atoi(l1[0])
    k, _ := strconv.Atoi(l1[1])

    var l2 []string

    for scanner.Scan() {
        l2 = strings.Fields(scanner.Text())
        if len(l2) > 0 {
            break
        }
    }

    if len(l2) < n {
        return 
    }

    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i], _ = strconv.Atoi(l2[i])
    }

    if backtrack(0, 0, k, n, arr) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }

}