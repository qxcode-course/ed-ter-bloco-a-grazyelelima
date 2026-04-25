package main

import "fmt"

func saida(v []int) string {
	saida := fmt.Sprintf("%v", v)
	if saida == "[]" {
		return "N"
	} else {
		return saida[1 : len(saida) - 1]
	}
}

func main() {

    var qtd_album, qtd_possui int
	fmt.Scan(&qtd_album, &qtd_possui)

	figuras := make([]int, qtd_possui)

	for i := range figuras {
		fmt.Scan(&figuras[i])
	}

	unicos := make(map[int]bool)
	repetidas := make([]int, 0, qtd_possui)

	for _, figura := range figuras {

		if unicos[figura] {
			repetidas = append(repetidas, figura)
		} else {
			unicos[figura] = true
		}
	}

	faltantes := make([]int , 0 , qtd_album)

	for i := 1; i <= qtd_album; i++ {
		if !unicos[i] {
			faltantes = append(faltantes, i)
		}
	}

	
	fmt.Println(saida(repetidas))
	fmt.Println(saida(faltantes))

}

