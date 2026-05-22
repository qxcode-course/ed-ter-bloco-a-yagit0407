package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func ToStr(l *list.List, sword *list.Element) string {
	var sb strings.Builder
	sb.WriteString("[")
	curr := l.Front()

	for curr != nil {
		val := curr.Value.(int)
		if curr == sword {
			if val > 0 {
				sb.WriteString(" " + strconv.Itoa(val) + ">")
			} else {
				sb.WriteString(" <" + strconv.Itoa(val))
			}
		} else {
			sb.WriteString(" " + strconv.Itoa(val))
		}
		curr = curr.Next()
	}
	sb.WriteString(" ]")
	return sb.String()
}

func Next(l *list.List, it *list.Element) *list.Element {
	if it.Next() == nil {
		return l.Front()
	}
	return it.Next()
}

func Prev(l *list.List, it *list.Element) *list.Element {
	if it.Prev() == nil {
		return l.Back()
	}
	return it.Prev()
}

func main() {
	var qtd, chosen, fase int
	fmt.Scan(&qtd, &chosen, &fase)
	l := list.New()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i * fase)
		fase = -fase
	}
	sword := l.Front()
	for i := 0; i < chosen-1; i++ {
		sword = Next(l, sword)
	}
	for i := 0; i < qtd-1; i++ {
		fmt.Println(ToStr(l, sword))
		if sword.Value.(int) > 0 {
			toRemove := Next(l, sword)
			sword = Next(l, toRemove)
			l.Remove(toRemove)
		} else {
			toRemove := Prev(l, sword)
			sword = Prev(l, toRemove)
			l.Remove(toRemove)
		}
	}
	fmt.Println(ToStr(l, sword))
}