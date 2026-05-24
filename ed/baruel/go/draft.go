package main

import "fmt"

func toString(vet[]int) string {
	saida := fmt.Sprintf("%v", vet)
	if saida == "[]" {
		return "N"
	} else {
		return (saida[1: len(saida) - 1])
	}
}

func main() {

	var qtdTotalFigurinhas, qtdFigurinhasDeBaruel int
	fmt.Scan(&qtdTotalFigurinhas, &qtdFigurinhasDeBaruel)

	figurinhas := make([]int, qtdFigurinhasDeBaruel)

	for i := range figurinhas {
		fmt.Scan(&figurinhas[i])
	}

	unicas := make(map[int]bool)
	repetidos := make([]int, 0, qtdFigurinhasDeBaruel)

	for _, figurinha := range figurinhas{
		if unicas[figurinha] {
			repetidos = append(repetidos, figurinha)

		} else {
			unicas[figurinha] = true
		}
	}

	faltando := make([]int, 0, qtdTotalFigurinhas)
	for i := 1; i <= qtdTotalFigurinhas; i++ {
		if !unicas[i] {
			faltando = append(faltando, i)
		}
	}

	fmt.Println(toString(repetidos))
	fmt.Println(toString(faltando))
	
}

