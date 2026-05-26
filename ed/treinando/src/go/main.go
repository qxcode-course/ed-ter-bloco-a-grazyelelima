package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	} //Se o vetor vier vazio, retorno "[]"

	var junta func(v []int) string

	junta = func(v []int) string {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}

		return strconv.Itoa(v[0]) + ", " + junta(v[1:])
	}

	return "[" + junta(vet) + "]"
}

func tostrrev(vet []int) string {
	
	if len(vet) == 0 {
		return "[]"
	}

	var juntaInvertido func(v []int) string

	juntaInvertido = func (v []int) string  {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}

		return juntaInvertido(v[1:]) + ", " + strconv.Itoa(v[0])
	}

	return "[" + juntaInvertido(vet) + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	if len(vet) <= 1 {
		return
	}//Se o meu vetor não tem nada ou só tem um elemento, ele já está invertido

	primeiroIndice := 0
	ultimoIndice := len(vet) - 1

	//Troca o primeiro pelo último e o último pelo primerio
	vet[primeiroIndice], vet[ultimoIndice] = vet[ultimoIndice], vet[primeiroIndice]

	reverse(vet[1:ultimoIndice]) //ignora as duas porntas

}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}

	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}

	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}


	var rec func(v []int) (int, int) 

	rec = func(v []int) (int, int) {
		if len(v) == 1 {
			return v[0], 0
		}

		indiceMenor, indice := rec(v[1:])
		indiceRealResto := indice + 1

		if v[0] < indiceMenor {
			return v[0], 0
		}

		return indiceMenor, indiceRealResto
	}

	_, indiceDeVerdade := rec(vet)

	return indiceDeVerdade
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
