package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct {
	data []int
	size int
	capacity int
}

func NewSet(capacity int) *Set {
	if capacity <= 0 {
		capacity = 4
	}

	return &Set{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (s *Set) reverse(newCapacity int) {
	if newCapacity <= s.capacity {
		return 
	}

	newData := make([]int, newCapacity)
	copy(newData, s.data[:s.size])
	s.data = newData
	s.capacity = newCapacity
}

func (s *Set) binarySearch(value int) int {
	low := 0
	high := s.size -1

	for low <= high {
		mid := low + (high-low)/2
		if s.data[mid] == value {
			return mid
		} else  if s.data[mid] < value {
			low = mid +1
		} else {
			high = mid -1
		}
	}

	return -(low +1)
}

func (s *Set) insert(value int, index int) error {
	if index < 0 || index > s.size {
		return fmt.Errorf("index out of bounds")
	}

	if s.size == s.capacity {
		if s.capacity == 0 {
			s.reverse(4) //Capacidade padrão
		} else {
			s.reverse(s.capacity * 2) //Dobra a capacidade
		}
	}

	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[index] = value
	s.size++

	return nil
}

func (s *Set) erase(index int) error {
	if index < 0 || index >= s.size {
		return fmt.Errorf("index out of bounds")
	}

	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}

	s.size--
	return nil
}

func (s *Set) Erase(value int) bool {
	idx := s.binarySearch(value)
	if idx < 0 {
		return false // Não encontrado
	}
	s.erase(idx)
	return true
}

func (s *Set) String() string {
	if s.size == 0 {
		return "[]"
	}
	strVals := make([]string, s.size)
	for i := 0; i < s.size; i++ {
		strVals[i] = strconv.Itoa(s.data[i])
	}
	return "[" + strings.Join(strVals, ", ") + "]"
}

func (s *Set) Insert(value int) {
	idx := s.binarySearch(value)
	if idx >= 0 {
		return // Se já existe, ignora!
	}
	// Decodifica a posição correta de inserção
	posInsercao := -idx - 1
	s.insert(value, posInsercao)
}

func (s *Set) Contains(value int) bool {
	return s.binarySearch(value) >= 0
}

func (s *Set) Clear() {
	s.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if !v.Contains(value) {
				fmt.Println("false")
			} else {
				fmt.Println("true")
			}
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
