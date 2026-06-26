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
	l := p.l
	c := p.c
	return []Pos {
		Pos{l:l + 0, c:c - 1},
		Pos{l:l + 0, c:c + 1},
		Pos{l:l  -1, c: c},
		Pos{l:l  +1, c: c + 0},
	}
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	nl := len(grid)
	nc := len(grid[0])

	if grid[l][c] == '#' {
		stack.Push(Pos{l: l, c: c})
	}

	for !stack.IsEmpty() {
		curr := stack.Pop()

		for _, visinho := range curr.getNeig() {
			if visinho.l >= 0 && visinho.l < nl && visinho.c >= 0 && visinho.c < nc {
				if grid[visinho.l][visinho.c] == '#' {
					grid[visinho.l][visinho.c] = 'o'
					stack.Push(visinho)
				}
			}
		}
	}


	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
