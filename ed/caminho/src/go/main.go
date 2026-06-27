package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{p.l + 1, p.c},
		{p.l, p.c + 1},
		{p.l - 1, p.c},
		{p.l, p.c - 1},
	}
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	nl := len(grid)
	nc := len(grid[0])

	visited := make([][]bool, nl)
	for i := range nl {
		visited[i] = make([]bool, nc)
	}

	parent := make(map[Pos]Pos)
	q := NewQueue[Pos]()
	q.Enqueue(startPos)
	visited[startPos.l][startPos.c] = true


	found := false
	for !q.IsEmpty() {
		curr, _ := q.Dequeue() 
		if curr == endPos {
			found = true
			break
		}
		for _, neig := range curr.getNeig() {
			if match(grid, neig, ' ') && !visited[neig.l][neig.c] {
				visited[neig.l][neig.c] = true
				parent[neig] = curr
				q.Enqueue(neig)
			}
		}
	}
	if found {
		grid[endPos.l][endPos.c] = '.'
		curr := parent[endPos]
		for curr != startPos {
			grid[curr.l][curr.c] = '.'
			curr = parent[curr]
		}
	}

	grid[startPos.l][startPos.c] = '.'
}

//func voltar()

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
