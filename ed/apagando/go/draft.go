package main

import "fmt"

func imprimir(filaFinal []int) {
	for i := 0; i < len(filaFinal); i++ {
		if i == len(filaFinal) -1 {
			fmt.Printf("%d ", filaFinal[i])
		} else {
			fmt.Printf("%d ", filaFinal[i])
		}
	}
	fmt.Println()
}

func main() {
	qtdFila := 0
	fmt.Scan(&qtdFila)

	fila := make([]int, qtdFila)
	for i := 0; i < qtdFila; i++ {
		fmt.Scan(&fila[i])
	}

	qtdSairam := 0
	fmt.Scan(&qtdSairam)

	sairam := make(map[int]bool)
	for i := 0; i < qtdSairam; i++ {
		var num int
		fmt.Scan(&num)
		sairam[num] = true
	}

	novaFila := make([]int, 0, qtdFila - qtdSairam)
	for i := 0; i < qtdFila; i++ {
		pessoa := fila[i]

		if !sairam[pessoa] {
			novaFila = append(novaFila, pessoa)
		}
	}

	imprimir(novaFila)


}
