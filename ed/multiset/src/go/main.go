package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Multiset struct {
	data []int
	size int
	capacity int
}

func NewMultiSet(capacity int) *Multiset {
	return &Multiset{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *Multiset) expand() {
	newCapacity := ms.capacity * 2
	if newCapacity == 0 {
		newCapacity = 4
	}

	newData := make([]int, newCapacity)
	copy(newData, ms.data[:ms.size])
	ms.data = newData
	ms.capacity = newCapacity
}

func (ms *Multiset) search(value int) (bool, int) {
	low := 0
	high := ms.size
	posInsercao := ms.size // Ponto padrão se for maior que todos

	for low < high {
		mid := low + (high-low)/2
		if ms.data[mid] < value {
			low = mid + 1
		} else {
			high = mid
		}
	}
	posInsercao = low

	// Agora verificamos se o elemento realmente existe na posição encontrada
	if posInsercao < ms.size && ms.data[posInsercao] == value {
		// Como queremos a última ocorrência, caminhamos para a direita enquanto houver iguais
		lastIdx := posInsercao
		for lastIdx+1 < ms.size && ms.data[lastIdx+1] == value {
			lastIdx++
		}
		return true, lastIdx
	}

	return false, posInsercao
}

func (ms *Multiset) insert(value int, index int) error {
	if index < 0 || index > ms.size {
		return fmt.Errorf("index out of bounds")
	}

	if ms.size == ms.capacity {
		ms.expand()
	}

	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[index] = value
	ms.size++
	return nil
}

func (ms *Multiset) erase(index int) error {
	if index < 0 || index >= ms.size {
		return fmt.Errorf("index out of bounds")
	}

	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}

	ms.size--
	return nil
}

func (ms *Multiset) Erase(value int) error {
	found, index := ms.search(value)
	if !found {
		return fmt.Errorf("value not found")
	}

	return ms.erase(index)
}
func (ms *Multiset) Insert(value int) {
	found, index := ms.search(value)

	if found {
		// Se já existe, o search nos deu a ÚLTIMA ocorrência. 
		// Inserimos em index + 1 para ele ficar logo após seus iguais.
		ms.insert(value, index+1)
	} else {
		// Se não existe, inserimos na posição de inserção retornada.
		ms.insert(value, index)
	}
}

func (ms *Multiset) Contains(value int) bool {
	found, _ := ms.search(value)
	return found
}

func (ms *Multiset) String() string {
	if ms.size == 0 {
		return "[]"
	}
	strVals := make([]string, ms.size)
	for i := 0; i < ms.size; i++ {
		strVals[i] = strconv.Itoa(ms.data[i])
	}
	return "[" + strings.Join(strVals, ", ") + "]"
}

func (ms *Multiset) Count(value int) int {
	found, index := ms.search(value)
	if !found {
		return 0
	}

	count := 0
	for i := index; i >= 0; i-- {
		if ms.data[i] == value {
			count++
		} else {
			break
		}
	}

	return count
}

func (ms *Multiset) Unique() int {
	if ms.size == 0 {
		return 0
	}

	uniqueCount := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			uniqueCount++
		}
	}
	return uniqueCount
}

func (ms *Multiset) Clear() {
	ms.size = 0 
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
			 	value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
