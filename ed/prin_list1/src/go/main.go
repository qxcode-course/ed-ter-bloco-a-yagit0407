package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	out := []string{}
	for node := l.Front(); node != l.End(); node = node.Next() {
		if node == sword {
			out = append(out, fmt.Sprintf("[%d]", node.Value))
		} else {
			out = append(out, fmt.Sprintf("%d", node.Value))
		}
	}
	return "[ " + strings.Join(out, " ") + " ]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it.Next() == l.End() {
		return l.Front()
	}
	return it.Next()
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for i := 0; i < chosen-1; i++ {
	sword = Next(l, sword)
	}
	for i := 0; i < qtd-1; i++ {
	fmt.Println(ToStr(l, sword))
	l.Erase(Next(l, sword))
	sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
	}
