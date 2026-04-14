package main

import "fmt"

func main() {
    var h, p, f, d int
   fmt.Scan(&h, &p, &f, &d) //Recebendo as entradas

   for {

        f += d //É aqui que a mágica acontece, irei alterar a posição do fugitivo de acordo com a direção, se for anti-horária eu vou somar, se for na direção horária eu vou diminuir, o for serve para isso, a cada interação eu vou diminuindo

        if f > 15 { //Como é de 0 a 15, ou seja, 16 números, eu vou fazer algumas verificações para o fugitivo sempre ficar nessa faixa de números
            f = 0
        } else if f < 0{
            f = 15
        }

        if f == h {
            fmt.Println("S")//Se eu encontrei o h, eu fugi, break encerra o laço for
            break
        }

        if f == p {
            fmt.Println("N")//Se eu for pego, dá o print e encerra o laço for com o break
            break
        }

   }
}    

