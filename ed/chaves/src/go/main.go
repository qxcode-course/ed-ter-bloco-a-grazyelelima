package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueue[string]()

	for i := 0; i < 16; i++ {
		queue.Enqueue(string(rune('A' + i)))
	}

	for i := 0; i < 15; i++ {
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		golsEsquerda, _ := strconv.Atoi(parts[0])
		golsDireita, _ := strconv.Atoi(parts[1])

		timeEsquerda := queue.Dequeue()
		timeDireita := queue.Dequeue()

		if golsEsquerda > golsDireita {
			queue.Enqueue(timeEsquerda)
		} else {
			queue.Enqueue(timeDireita)
		}
	}

	fmt.Println(queue.Dequeue())
}
