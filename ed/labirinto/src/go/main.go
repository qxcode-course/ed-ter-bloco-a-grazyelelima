package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func getNeig(p Pos) []Pos {
	return []Pos{
			{p.l - 1, p.c}, //Cima
			{p.l + 1, p.c}, //Baixo
			{p.l, p.c - 1}, //Esqueda
			{p.l, p.c + 1}, //Direita
    }
}

func inside(grid [][]rune, p Pos) bool {
	nl := len(grid)
	nc := len(grid[0])
	
	if p.c < 0 || p.c >= nc || p.l < 0 || p.l >= nl {
		return false
	} 

	return true
}

func match(grid [][]rune, p Pos, value rune) bool {
	return inside(grid, p) && grid[p.l][p.c] == value //Verifico se a posição é válida e se o caractere bate com o valor esperado
}

// Função recursiva que tenta encontrar o caminho do início ao fim
func search(grid [][]rune, startPos, endPos Pos, visited[][]bool) bool {
	visited[startPos.l][startPos.c] = true //Marca o caminho como visitado para não andar em círculo
	grid[startPos.l][startPos.c] = '.'
	if startPos == endPos {
		return true //Se a posição inicial for igual a posição final, retorne true
	}

	neighbors := getNeig(startPos) // pega os 4 vizinhos

	for _, nextPos := range neighbors {

		if match(grid, nextPos, ' ') && !visited[nextPos.l][nextPos.c] {

			if search(grid, nextPos,  endPos, visited) {
				return true
			}
		}
	}
	grid[startPos.l][startPos.c] = ' '

	return false

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	visited := make([][]bool, nl)
	for i := range visited {
		visited[i] = make([]bool, nc)
	}

	search(grid, startPos, endPos, visited)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
