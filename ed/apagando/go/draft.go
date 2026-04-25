package main

import "fmt"

func main() {

	var n, m int 

	fmt.Scan(&n)

	fila := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&m)

	quemSai := make(map[int]bool)

	for i := 0; i < m; i++ {
		var id int
		fmt.Scan(&id)
		quemSai[id] = true
 	}	

	primeiro := true

	for _, pessoa := range fila {
		if !quemSai[pessoa] {
			if !primeiro {
				fmt.Print(" ")
			}
			fmt.Print(pessoa)
			primeiro = false
		}
	}
	fmt.Println(" ")
}
