package main

import  "fmt" 

type Competidor struct {
	pedraA int
	pedraB int
}

func difAbs(c Competidor) int {
	if c.pedraA > c.pedraB {
		return c.pedraA - c.pedraB
	}
	return c.pedraB - c.pedraA
}

func main() {
    nCompetidores := 0
	fmt.Scan(&nCompetidores)
	
	jogadas := make([]Competidor, nCompetidores)
	
	for i := range jogadas {
		fmt.Scan(&jogadas[i].pedraA, &jogadas[i].pedraB)
	}

	melhorJogadaIndice := -1

	for i, jog := range jogadas {
		if jog.pedraA < 10 || jog.pedraB < 10 {
			continue
		}

		if melhorJogadaIndice == -1 || difAbs(jog) < difAbs(jogadas[melhorJogadaIndice]) {
			melhorJogadaIndice = i
		}

	}

	if melhorJogadaIndice == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(melhorJogadaIndice)
	}
}
