package main
import "fmt"
func main() {
    var qtd_figurinhasAlbum int
    var qtd_figurinhasBoruel int

    fmt.Scan(&qtd_figurinhasAlbum)
    fmt.Scan(&qtd_figurinhasBoruel)

    tem := make([]int, qtd_figurinhasBoruel)
    existeNoAlbum := make([]bool, qtd_figurinhasAlbum+1)
    var repetidas []int

    for i := 0; i < qtd_figurinhasBoruel; i++ {
		fmt.Scan(&tem[i])
		
		if i > 0 && tem[i] == tem[i-1] {
			repetidas = append(repetidas, tem[i])
		}
		
		if tem[i] <= qtd_figurinhasAlbum {
			existeNoAlbum[tem[i]] = true
		}
	}

    if len(repetidas) == 0 {
		fmt.Println("N")
	} else {
		for i, v := range repetidas {
			fmt.Print(v)
			if i < len(repetidas)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	houveFaltante := false
	primeiroFaltante := true
	for i := 1; i <= qtd_figurinhasAlbum; i++ {
		if !existeNoAlbum[i] {
			if !primeiroFaltante {
				fmt.Print(" ")
			}
			fmt.Print(i)
			houveFaltante = true
			primeiroFaltante = false
		}
	}

	if !houveFaltante {
		fmt.Print("N")
	}
	fmt.Println()
}

