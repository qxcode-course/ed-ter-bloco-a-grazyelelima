package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	if len(grid) == 0 || len(word) == 0 {
		return false
	}

	linhas := len(grid)
	colunas := len(grid[0])

	for r := 0; r < linhas; r++ {
		for c := 0; c < colunas; c++ {
			if grid[r][c] == word[0] {
				if dfs(r, c, 0, grid, word) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(r int, c int, index int, grid [][]byte, word string) bool {
	if index == len(word) {
		return true
	}

	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != word[index] {
		return false
	}

	t := grid[r][c]
	grid[r][c] = '#'

	found := dfs(r+1, c, index+1, grid, word) || 
			dfs(r-1, c, index+1, grid, word) ||
			dfs(r, c+1, index+1, grid, word) ||
			dfs(r, c-1, index+1, grid, word)

	grid[r][c] = t
	return found
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
