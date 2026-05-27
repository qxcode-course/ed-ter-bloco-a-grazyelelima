package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func getMen(vet []int) []int {
	var homens []int

	for _, pessoa := range vet {
		if pessoa > 0 { //Se o número for positivo, é homem
			homens = append(homens, pessoa)
		}
	}

	return homens //Devolve uma lista com os homens encontrados
}

func getCalmWomen(vet []int) []int {
	var mulheresCalmas []int

	for _, pessoa := range vet {
		if pessoa < 0 && pessoa >= -9 { //Se o número for negativo, é mulher e são calmas se o estresse for inferior a 10, ou seja, vai de -1 até -9
			mulheresCalmas = append(mulheresCalmas, pessoa)
		}
	}

	return mulheresCalmas //retorno uma lista com as mulheres encontradas
}


func sortVet(vet []int) []int {
	copia := append([]int{}, vet...) //cria uma cópia do original sem alterar sua ordem na main
	n := len(vet) //Amarzeno o tamanho, ou seja, a quantidade de elementos

	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {
			if copia[j] > copia[j+1] {
				copia[j], copia[j+1] = copia[j+1], copia[j]
			}
		}
	}

	return copia

}

func sortStress(vet []int) []int {
	copia := append([]int{}, vet...) //cria uma cópia do original sem alterar sua ordem na main
	n := len(vet) //Amarzeno o tamanho, ou seja, a quantidade de elementos

	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {
			estresseEsqueda := math.Abs(float64(copia[j]))
			estresseDireita := math.Abs(float64(copia[j+1]))

			if estresseEsqueda > estresseDireita {
				copia[j], copia[j+1] = copia[j+1], copia[j]
			}
		}
	}

	return copia
}

func reverse(vet []int) []int {
	var invertido []int //vetor que vai guardar os valores invertidos

	for i := len(vet) - 1; i >= 0; i-- { //percorre o vetor do último elemento até o primeiro
		invertido = append(invertido, vet[i]) //vai preenchendo o vetor invertido do ultimo elemento do vet até o primeiro
	}

	return invertido
}

func unique(vet []int) []int {
	var unico []int //vetor para guardar os calores que não se repetem

	for i := 0; i < len(vet); i++ {
		jaExiste := false
		for j := 0; j < len(unico); j++ {
			if vet[i] == unico[j] {
				jaExiste = true
				break
			}
		}

		if !jaExiste {
			unico = append(unico, vet[i])
		}
	} 

	return unico
}

func repeated(vet []int) []int {
	var seRepete []int

	for i := 0; i < len(vet); i++ {
		isCopia := false
		for j := 0; j < i; j++ {
			if vet[i] == vet[j] {
				isCopia = true
				break
			}
		}

		if isCopia {
			jaSalvo := false
			for l := 0; l < len(seRepete); l++ {
				if vet[i] == seRepete[l] {
					jaSalvo = true
					break
				}
			}

			if !jaSalvo || (len(vet) > 0 && vet[0] == vet[len(vet)-1] && len(seRepete) < len(vet)-1) {
				seRepete = append(seRepete, vet[i])
			}
		}
	}
	return seRepete
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

