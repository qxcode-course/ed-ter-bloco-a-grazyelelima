package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data  []int
	index int
}

type ReverseIterador struct {
	data []int
	index int
}

type CyclicIterador struct {
	data []int
	index int
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: -1}
}


func (l *MyList) ReverseIterador() *ReverseIterador {
	return &ReverseIterador{data: l.data, index: len(l.data)}
}

func (l *MyList) CyclicIterador() *CyclicIterador {
	return &CyclicIterador{data: l.data, index: -1}
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)-1
}

func (r *ReverseIterador) HasNext() bool {
	return r.index > 0
}

func (c *CyclicIterador) HasNext() bool {
	return len(c.data) > 0
}

func (i *Iterator) Next() int {
	if i.index == len(i.data) {
		panic(fmt.Errorf("No more elements"))
	}
	i.index += 1
	return i.data[i.index]
}

func (r *ReverseIterador) Next() int {
	if r.index <= 0 {
		panic(fmt.Errorf("No more elements"))
	}
	r.index--
	return r.data[r.index]
}

func (c *CyclicIterador) Next() int {
	if len(c.data) == 0 {
		panic(fmt.Errorf("No elements in cyclic list"))
	}

	c.index++
	return c.data[c.index % len(c.data)]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			break
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					slice[i], _ = strconv.Atoi(value)
				}
				mylist = NewMyList(slice)
			}
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterador(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "cyclic":
			qtd, _ := strconv.Atoi(args[1])
			fmt.Print("[ ")
			it := mylist.CyclicIterador()
			for range qtd {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		}
	}

}
