package main

import "fmt"


func acharEnesimo(n, atual int) int {
    if eh_primo(atual, 2) {
        if n == 1 {
            return atual
        }
        return acharEnesimo(n-1, atual+1)
    }
    return acharEnesimo(n, atual+1)
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
	fmt.Println(acharEnesimo(x, 2))
}
