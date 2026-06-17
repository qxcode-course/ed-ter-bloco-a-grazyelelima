package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type LList struct {
	root *Node
	size int
}

type Node struct {
	Value int
	next *Node
	prev *Node
}

func NewLList() LList {
	sentinel := &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return LList {
		root: sentinel,
		size: 0,
	}
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) insert(value int, anterior *Node, proximo *Node) {
	novo := &Node{Value: value, prev: anterior, next: proximo}
	anterior.next = novo
	proximo.prev = novo
	ll.size++
}

func (ll *LList) remove(alvo *Node) {
	if alvo == ll.root {
		return
	}

	anterior := alvo.prev
	proximo := alvo.next

	anterior.next = proximo
	proximo.prev = anterior
	ll.size--
}

func (ll *LList) PushFront(value int) {
	ll.insert(value, ll.root, ll.root.next)
}

func (ll *LList) PushBack(value int) {
	ll.insert(value, ll.root.prev, ll.root)
}

func (ll *LList) PopFront() {
	if ll.size > 0 {
		ll.remove(ll.root.next)
	}
}

func (ll *LList) PopBack() {
	if ll.size > 0 {
		ll.remove(ll.root.prev)
	}
}

func (ll *LList) String() string {
	if ll.size == 0 {
		return "[]"
	}

	var elementos []string
	node := ll.root.next
	for node != ll.root {
		elementos = append(elementos, strconv.Itoa(node.Value))
		node = node.next
	}
	return "[" + strings.Join(elementos, ", ") + "]"
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
