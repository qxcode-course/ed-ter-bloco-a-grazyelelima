package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack[T any] struct {
    data []T
}

func NewStrack[T any]() *Stack[T] {
    return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Push(value T) {
    s.data = append(s.data, value)
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.data) == 0
}

func (s *Stack[T]) Top() T {
    if s.IsEmpty() {
        var zero T
        return zero
    }
    return s.data[len(s.data)-1]
}

func (s *Stack[T]) Pop() T {
    if s.IsEmpty() {
        var zero T
        return zero
    }

    value := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return value
}

func IsBalanced(str string) bool {
    stack := NewStrack[rune]()

    for _, char := range str {
        if char == '(' || char == '[' {
            stack.Push(char)
        } else if char == ')' || char == ']' {
            if stack.IsEmpty() {
                return false
            }

            topo := stack.Top()

            if char == ')' && topo == '(' {
                stack.Pop()
            } else if char == ']' && topo == '[' {
                stack.Pop()
            } else {
                return false
            }
        }
    }
    return stack.IsEmpty()
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan() {
        line := scanner.Text()

        if IsBalanced(line) {
            fmt.Println("balanceado")
        } else {
            fmt.Println("nao balanceado")
        }
    }
}