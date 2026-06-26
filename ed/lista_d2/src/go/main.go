package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)


//estrutura do node

type Node struct {
	Value int
	next *Node
	prev *Node
	root *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}

	return n.prev
}

//estrutura da lista:

type LList struct {
	root *Node
	size int
}

func NewLList() LList {
	sentinel := &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	sentinel.root = sentinel
	return LList{
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

//metodo auxiliar que vai evitar repetições no cógigo, top dms
func (ll *LList) insertAux(value int, proximo *Node) *Node {
	anterior := proximo.prev
	novo := &Node{Value: value, prev: anterior, next: proximo, root: ll.root}
	anterior.next = novo
	proximo.prev = novo
	ll.size++
	return novo
}

func (ll *LList) PushFront(value int) {
	ll.insertAux(value, ll.root.next)
}

func (ll *LList) PushBack(value int) {
	ll.insertAux(value, ll.root)
}

func (ll *LList) Remove(node *Node) *Node {
	if node == nil || node == ll.root {
		return nil
	}
	anterior := node.prev
	proximo := node.next
	anterior.next = proximo
	proximo.prev = anterior

	ll.size--

	if proximo == ll.root {
		return nil
	}
	return proximo
}

func (ll *LList) PopFront() {
	if ll.size > 0 {
		ll.Remove(ll.root.next)
	}
}

func (ll *LList) PopBack() {
	if ll.size > 0 {
		ll.Remove(ll.root.prev)
	}
}

func (ll *LList) Front() *Node {
	if ll.size == 0 {
		return nil
	}

	return ll.root.next
}

func (ll *LList) Back() *Node {
	if ll.size == 0 {
		return nil
	}

	return ll.root.prev
}

func (ll *LList) Search(value int) *Node {
	for node := ll.root.next; node != ll.root; node = node.next {
		if node.Value == value {
			return node
		}
	}
	return nil
}

func (ll *LList) Insert(node *Node, value int) {
	if node == nil {
		return
	}
	ll.insertAux(value, node)
}

func (ll *LList) String() string {
	if ll.size == 0 {
		return "[]"
	}

	var elements []string
	for node := ll.Front(); node != nil; node = node.Next() {
		elements = append(elements, strconv.Itoa(node.Value))
	}

	return "[" + strings.Join(elements, ", ") + "]"
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
			 	fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.Value)
			}
			 fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
			 	node.Value = newvalue
			} else {
			 	fmt.Println("fail: not found")
			 }
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Insert(node, newvalue)
			} else {
			 	fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Remove(node)
			} else {
			 	fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
