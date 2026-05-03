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
	Next *Node
	Prev *Node
}

type LList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLList() *LList {
	return &LList{}
}

func (l *LList) PushBack(value int) {
	
	node := &Node {
		Value : value,
	}
	if l.Size == 0 {
		
		l.Head = node
		l.Tail = node
	} else {
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}
	l.Size++
}

func (l *LList) String() string {

	if l.Size == 0 {
		return "[]"
	}

	text := "["

	aux := l.Head

	for aux != nil {

		text += fmt.Sprintf("%d", aux.Value)

		if aux.Next != nil {
			text += ", "
		}

		aux = aux.Next
	}

	text += "]"

	return text
}

func (l *LList) Clear() {

	l.Head = nil
	l.Tail = nil

	l.Size = 0
}

func (l *LList) PushFront(value int) {

	node := &Node{
		Value: value,
	}

	if l.Size == 0 {

		l.Head = node
		l.Tail = node

	} else {

		node.Next = l.Head

		l.Head.Prev = node

		l.Head = node
	}

	l.Size++
}

func (l *LList) PopBack() {

	if l.Size == 0 {
		
		return
	}

	if l.Size == 1 {

		l.Head = nil
		l.Tail = nil

	} else {

		l.Tail = l.Tail.Prev

		l.Tail.Next = nil
	}

	l.Size--
}

func (l *LList) PopFront() {

	if l.Size == 0 {
		
		return
	}

	if l.Size == 1 {

		l.Head = nil
		l.Tail = nil

	} else {

		l.Head = l.Head.Next

		l.Head.Prev = nil
	}

	l.Size--
}

func (l *LList) SizeList() int {
	return l.Size
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
			fmt.Println(ll.SizeList())
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
