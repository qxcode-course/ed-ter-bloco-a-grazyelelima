package main

import  "fmt" 

type Jogada struct {
	a, b int
}

func valorAbs(j Jogada) int {
	if j.a > j.b {
		return j.a - j.b
	}

	return j.b - j.a
}

func main() {
    
	n := 0
	fmt.Scan(&n)

	jogadas := make([]Jogada, n)

	for i := range n {
		fmt.Scan(&jogadas[i].a, &jogadas[i].b)
	}

	indice_melhor := -1
	for i, jog := range jogadas {
		if jog.a < 10 || jog.b < 10 {
			continue
		}

		if indice_melhor == -1  || valorAbs(jog) < valorAbs(jogadas[indice_melhor]){
			indice_melhor = i
		}
	}

	if indice_melhor == -1 {
			fmt.Println("sem ganhador")
		} else {
			fmt.Println(indice_melhor)
		}
}
