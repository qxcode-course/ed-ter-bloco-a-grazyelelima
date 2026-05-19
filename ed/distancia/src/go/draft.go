package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(s string, L int) string {
	n := len(s)
	period := L + 1
	cycle := make([]int, period)
	for i := 0; i < period; i++ {
		cycle[i] = -1
	}
	used := make([]bool, L+1)

	// Registra os dígitos fixos no ciclo
	for i, ch := range s {
		if ch != '.' {
			d := int(ch - '0')
			r := i % period
			if cycle[r] == -1 {
				cycle[r] = d
				used[d] = true
			}
			// Se houver conflito, a entrada seria inválida, mas o problema garante que não ocorre
		}
	}

	// Coleta os dígitos que não apareceram nos fixos
	missing := []int{}
	for d := 0; d <= L; d++ {
		if !used[d] {
			missing = append(missing, d)
		}
	}

	// Preenche as posições vazias do ciclo com os dígitos faltantes
	idx := 0
	for r := 0; r < period; r++ {
		if cycle[r] == -1 {
			cycle[r] = missing[idx]
			idx++
		}
	}

	// Gera a string resultado
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = byte('0' + cycle[i%period])
	}
	return string(res)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	L, _ := strconv.Atoi(scanner.Text())
	fmt.Println(solve(s, L))
}