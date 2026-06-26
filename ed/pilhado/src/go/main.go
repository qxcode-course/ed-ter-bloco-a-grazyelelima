package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	l, c int
}

func main () {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return 
	}

	dimencoes := strings.Fields(scanner.Text())
	nl, _ := strconv.Atoi(dimencoes[0])
	nc, _ := strconv.Atoi(dimencoes[1])

	var inicio, fim Pos

	grid := make([][]rune, nl)
	for i := 0; i < nl; i++ {
		if scanner.Scan() {
			grid[i] = []rune(scanner.Text())

			if len(grid[i]) < nc {
				linhaCompleta := make([]rune, nc)
				copy(linhaCompleta, grid[i])
				for j := len(grid[i]); j < nc; j++ {
					linhaCompleta[j] = ' '
				}
				grid[i] = linhaCompleta
			}

			for j := 0; j < nc; j++ {
				if grid[i][j] == 'I' {
					inicio = Pos{l: i, c: j}
				} else if grid[i][j] == 'F' {
					fim = Pos{l: i, c: j}
				}
			}
		}
	}


	caminho := NewStack[Pos]()
	becos := NewStack[Pos]()
	visited := make(map[Pos]bool)

	caminho.Push(inicio)

	dl := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	for !caminho.IsEmpty() {
		atual := caminho.Top()
		visited[atual] = true

		if atual == fim {
			break
		}

		var validos []Pos

		for i := 0; i < 4; i++ {
			nlViz := atual.l + dl[i]
			ncViz := atual.c + dc[i]

			if nlViz >= 0 && nlViz < nl && ncViz >=0 && ncViz < nc {
				vizinho := Pos{l: nlViz, c: ncViz}
				if grid[nlViz][ncViz] != '#' && !visited[vizinho] {
					validos = append(validos, vizinho)
				}
			}
		}

		if len(validos) > 0 {
			caminho.Push(validos[0])
		} else {
			becos.Push(atual)
			caminho.Pop()
		}
	}

	for !caminho.IsEmpty() {
		pos := caminho.Pop()
		grid[pos.l][pos.c] = '.'
	}

	for !becos.IsEmpty() {
		beco := becos.Pop()
		grid[beco.l][beco.c] = ' '
	}

	for i := 0; i < nl; i++ {
		fmt.Println(string(grid[i]))
	}

}

