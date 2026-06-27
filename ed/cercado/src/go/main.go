package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	nl := len(board)
	nc := len(board[0])

	if nl == 0 || nc == 0 {
		return 
	}

	for i := 0; i < nl; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][nc-1] == 'O' {
			dfs(board, i, nc-1)
		}
	}

	for j := 0; j < nc; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}
		if board[nl-1][j] == 'O' {
			dfs(board, nl-1, j)
		}
	}

	for k := 0; k < nl; k++ {
		for l := 0; l < nc; l++ {
			if board[k][l] == 'O' {
				board[k][l] = 'X'
			} else if board[k][l] == 'E' {
				board[k][l] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, l, c int) {
	nl := len(board)
	nc := len(board[0])

	if l < 0 || c < 0 || l >= nl || c >= nc || board[l][c] != 'O' {
		return
	}

	board[l][c] = 'E'

	dfs(board, l+1, c)
	dfs(board, l-1, c)
	dfs(board, l, c+1)
	dfs(board, l, c-1)
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
