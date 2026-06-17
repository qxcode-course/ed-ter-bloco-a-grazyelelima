package main

import (
	"bufio"
	"fmt"
	"strings"
	"errors"
	"strconv"
	"os"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}

func (v *Vector) String() string {
	if v.size == 0 {
		return "[]"
	}
	
	elem := Join(v.data[:v.size], ", ")
	return "[" + elem + "]"
}

func (v *Vector) PushBack(value int) {
    if v.size == v.capacity {
        novaCap := 0
        if v.capacity == 0 {
            novaCap = 1
        } else {
            novaCap = v.capacity * 2
        }

        novoData := make([]int, novaCap)

        for i := 0; i < v.size; i++ {
            novoData[i] = v.data[i]
        }
    
        v.data = novoData
        v.capacity = novaCap
    }

    v.data[v.size] = value
    v.size = v.size + 1
}

func (v *Vector) At(index int) (int, error) {
	if index < 0 || index >= v.size {
        return 0, errors.New("index out of range")
    }
    return v.data[index], nil
}

func (v *Vector) Set(index int, value int) error {
    if index < 0 || index >= v.size {
        return errors.New("index out of range")
    }
    v.data[index] = value
    return nil
}

func (v *Vector) Clear() {
	v.size = 0
}

func (v *Vector) Reserve(newCapacity int) {
	if newCapacity > v.capacity {
		novoData := make([]int, newCapacity)
		for i := 0; i < v.size; i++ {
			novoData[i] = v.data[i]
		}
		v.data = novoData
		v.capacity = newCapacity
	}
}

func (v *Vector) PopBack() error {
	if v.size == 0 {
        return errors.New("vector is empty")
    }
    v.size = v.size - 1
    return nil
}

func (v *Vector) Insert(index int, value int) error {
	if index < 0 || index > v.size {
        return errors.New("index out of range")
    }
    if v.size == v.capacity {
        novaCap := 1
        if v.capacity > 0 {
            novaCap = v.capacity * 2
        }
        novoData := make([]int, novaCap)
        for i := 0; i < index; i++ {
            novoData[i] = v.data[i]
        }
        for i := index; i < v.size; i++ {
            novoData[i+1] = v.data[i]
        }
        v.data = novoData
        v.capacity = novaCap
    } else {
        for i := v.size; i > index; i-- {
            v.data[i] = v.data[i-1]
        }
    }
    v.data[index] = value
    v.size = v.size + 1
    return nil
}

func (v *Vector) Erase(index int) error {
	if index < 0 || index >= v.size {
		return errors.New("index out of range")
	}

	for i := index; i < v.size -1; i++ {
		v.data[i] = v.data[i+1]
	} 

	v.size = v.size -1
	return nil
}

func (v *Vector) IndexOf(value int) int {
	index := -1
	for i := 0; i < v.size-1; i++ {
		if  v.data[i] == value {
			index = i
			break
		}
	}

	return index
}

func (v * Vector) Contains(value int) bool {
	if v.IndexOf(value) != -1 {
		return true
	}

	return false
}

func (v *Vector) Slice(start, end int) * Vector {
	if end == -1 {
		end = v.size -1
	}

	if start < 0 || start > v.size {
		start = v.size
	}

	if end < 0 || end > v.size {
		end = v.size
	}

	if start > end {
		start = end
	}

	length := end - start
	novo := &Vector{
		data: make([]int, length),
		size: length,
		capacity: length,
	}

	for i := 0; i < length; i++ {
		novo.data[i] = v.data[start+i]
	}

	return novo
}

func (v *Vector) Capacity() int {
	return v.capacity
}


func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			} 
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
