package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	Value int
	prev *Node
	next *Node
}

type LList struct {
	head *Node
	tail *Node
	size int
}

func NewLList() *LList {
	return &LList{}
}

func (l *LList) PushBack(value int) {
	node := &Node{Value : value}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}

	l.size++
}

func (l *LList) PushFront (value int) {
	node := &Node{Value : value}

	if l.head == nil {
		l.head = node 
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.size++
}


func (l *LList) String() string {
	out := "["

	for node := l.head; node != nil; node = node.next {

		out += fmt.Sprintf("%d", node.Value)

		if node.next != nil {
			out += ", "
		}
	}

	out += "]"

	return out
}

func (l *LList) PopBack() {
	if l.tail == nil {
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}

	l.size--
}

func (l *LList) PopFront() {
	if l.head == nil {
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}

	l.size--
}

func (l *LList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) Search(value int) *Node {
	for node := l.head; node != nil; node = node.next {
		if node.Value == value {
			return node
		}

	}
	return nil
}

func (l *LList) Insert(node *Node, value int) {

	if node == l.head {
		l.PushFront(value)
		return
	}

	newNode := &Node{Value: value}

	prev := node.prev

	prev.next = newNode
	newNode.prev = prev

	newNode.next = node
	node.prev = newNode

	l.size++
}

func (l *LList) Remove(node *Node) {

	if node == nil {
		return
	}

	
	if node == l.head {
		l.PopFront()
		return
	}

	
	if node == l.tail {
		l.PopBack()
		return
	}

	
	node.prev.next = node.next
	node.next.prev = node.prev

	l.size--
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

			for node := ll.head; node != nil; node = node.next {
				fmt.Printf("%v ", node.Value)
			}

			fmt.Print("]\n[ ")

			for node := ll.tail; node != nil; node = node.prev {
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
