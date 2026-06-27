package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Não modifique a assinatura da função floodFill
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	oldColor := image[sr][sc]

	if oldColor == color {
		return image
	}

	dfs(image, sr, sc, oldColor, color)
	return image
}

func dfs(image [][]int, l, c int, oldColor, newColor int) {
	nl := len(image)
	nc := len(image[0])

	if l < 0 || c < 0 || l >= nl || c >= nc || image[l][c] != oldColor {
		return
	}

	image[l][c] = newColor

	dfs(image, l+1, c, oldColor, newColor)
	dfs(image, l-1, c, oldColor, newColor)
	dfs(image, l, c+1, oldColor, newColor)
	dfs(image, l, c-1, oldColor, newColor)
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	// Lê sr, sc e color
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])

	result := floodFill(image, sr, sc, color)

	// Imprime a matriz resultante
	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
