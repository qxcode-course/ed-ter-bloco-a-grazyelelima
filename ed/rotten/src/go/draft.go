package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue[T any] struct {
    queue *list.List
}

func NewQueue[T any]() *Queue[T] {
    return &Queue[T]{queue: list.New()}
}
func (q *Queue[T]) Enqueue(value T) {
    q.queue.PushBack(value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    element := q.queue.Front()
    if element == nil {
        var zero T
        return zero, false
    }
    q.queue.Remove(element)
    return element.Value.(T), true
}

func (q *Queue[T]) IsEmpty() bool {
    return q.queue.Len() == 0
}

type Pos struct {
    r, c int
}

func orangesRotting(grid [][]int) int {
    nrows := len(grid)
    ncols := len(grid[0])

    q := NewQueue[Pos]()
    frescas := 0
    for r := 0; r < nrows; r++ {
        for c := 0; c < ncols; c++ {
            if grid[r][c] == 2 {
                q.Enqueue(Pos{r, c})
            } else if grid[r][c] == 1 {
                frescas++
            }
        }
    }
    if frescas == 0 {
        return 0
    }

    minutos := 0
    dr := []int{-1, 1, 0, 0}
    dc := []int{0, 0, -1, 1}

    for !q.IsEmpty() {
        tamanhoOnda := q.queue.Len()
        algumaContaminada := false

        for i := 0; i < tamanhoOnda; i++ {
            curr, _ := q.Dequeue()

            for j := 0; j < 4; j++ {
                nr := curr.r + dr[j]
                nc := curr.c + dc[j]

                if nr >= 0 && nc >= 0 && nr < nrows && nc < ncols && grid[nr][nc] == 1 {
                    grid[nr][nc] = 2
                    frescas--
                    algumaContaminada = true
                    q.Enqueue(Pos{nr, nc})
                }
            }
        }
        if algumaContaminada {
            minutos++
        }
    }

    if frescas > 0 {
        return -1
    }
    return minutos
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }
    var nrows, ncols int
    fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
    grid := make([][]int, nrows)
    for i := 0; i < nrows; i++ {
        scanner.Scan()
        fields := strings.Fields(scanner.Text())
        grid[i] = make([]int, ncols)
        for j := 0; j < ncols; j++ {
            grid[i][j], _ = strconv.Atoi(fields[j])
        }
    }
    fmt.Println(orangesRotting(grid))
}