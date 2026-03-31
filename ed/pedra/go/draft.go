package main

import  "fmt" 

type Jogada struct {
    pa, pb int
}

func pontuacao(j Jogada) int {
    if j.pa > j.pb {
        return j.pa - j.pb
    }
    return j.pb - j.pa
}

func main() {
    n := 0

    fmt.Scan(&n)

    jogadas := make([]Jogada, n)

    for i := range n {
        fmt.Scan(&jogadas[i].pa, &jogadas[i].pb)
    }

    ind_melhor := -1

    for i, jog := range jogadas {
        if jog.pa < 10 || jog.pb < 10 {
            continue
        } 

        if ind_melhor == -1 || (pontuacao(jog) < pontuacao(jogadas[ind_melhor])) {
            ind_melhor = i
        }
    }

    if ind_melhor == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(ind_melhor)
    }



}
