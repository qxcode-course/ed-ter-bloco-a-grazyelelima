package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EhVazio(vet []int) bool {
	return fmt.Sprintf("%v", vet) == "[]" //Transformo em string e verifico se é igual a "[]", significa que é vazio
}

// não altere a assinatura
func equals(a []int, b []int) bool {
	if EhVazio(a) && EhVazio(b) {
		return true //Se ambos são vazios, eles são iguais
	}

	if EhVazio(a) || EhVazio(b) {
		return false //Se só um dos dois é vazio, então retorna falso, eles não são iguais
	}

	if a[0] != b[0] {
		return false //Se o primeiro elemento do vetor a for diferente do primeiro elemento do vetor b, retorna falso, não são iguais
	}

	return equals(a[1:], b[1:])//Chamada recursiva, compara se os primeiros elementos são iguais, dps corta e o próximo elemento vira o primeiro e assim por diante, até o vetor ficar vazio...caso seja igual

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	vet1 := str2slice(scanner.Text())
	scanner.Scan()
	vet2 := str2slice(scanner.Text())
	if equals(vet1, vet2) {
		fmt.Println("iguais")
	} else {
		fmt.Println("diferentes")
	}
}

func str2slice(line string) []int {
	parts := strings.Fields(line)
	nums := make([]int, 0)
	for i := 1; i < len(parts)-1; i++ {
		value, _ := strconv.Atoi(parts[i])
		nums = append(nums, value)
	}
	return nums
}
