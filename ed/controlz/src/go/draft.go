package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type OpType int

const (
	Insert OpType = iota
	Delete
)

type Action struct {
	TypeOp OpType
	value  rune
	Before *list.Element
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	line := scanner.Text()

	l := list.New()
	cursor := l.PushBack('|')

	var undoStack []Action
	var redoStack []Action // Declarando a pilha de refazer

	pushToUndo := func(act Action) {
		undoStack = append(undoStack, act)
		redoStack = nil
	}

	for _, ch := range line {
		switch ch {
		case '<':
			if cursor.Prev() != nil {
				prev := cursor.Prev()
				l.Remove(cursor)
				cursor = l.InsertBefore('|', prev)
			}
		case '>':
			if cursor.Next() != nil {
				next := cursor.Next()
				l.Remove(cursor)
				cursor = l.InsertAfter('|', next)
			}
		case 'B':
			if cursor.Prev() != nil {
				toRemove := cursor.Prev()
				val := toRemove.Value.(rune)
				before := toRemove.Prev()
				l.Remove(toRemove)
				pushToUndo(Action{TypeOp: Delete, value: val, Before: before})
			}
		case 'D':
			if cursor.Next() != nil {
				toRemove := cursor.Next()
				val := toRemove.Value.(rune)
				before := cursor.Prev()
				l.Remove(toRemove)
				pushToUndo(Action{TypeOp: Delete, value: val, Before: before})
			}
		case 'R', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '-':
			val := ch
			if ch == 'R' {
				val = '\n'
			}
			inserted := l.InsertBefore(val, cursor)
			pushToUndo(Action{TypeOp: Insert, value: val, Before: inserted.Prev()})
		case 'Z':
			if len(undoStack) > 0 {
				act := undoStack[len(undoStack)-1]
				undoStack = undoStack[:len(undoStack)-1]
				redoStack = append(redoStack, act)

				if act.TypeOp == Insert {
					curr := l.Front()
					for curr != nil {
						if curr != cursor && curr.Value.(rune) == act.value && curr.Prev() == act.Before {
							l.Remove(curr)
							break
						}
						curr = curr.Next()
					}
				} else if act.TypeOp == Delete {
					if act.Before == nil {
						l.PushFront(act.value)
					} else {
						l.InsertAfter(act.value, act.Before)
					}
				}
			}
		case 'Y': // Agora posicionado corretamente no switch externo
			if len(redoStack) > 0 {
				act := redoStack[len(redoStack)-1]
				redoStack = redoStack[:len(redoStack)-1]
				undoStack = append(undoStack, act)

				if act.TypeOp == Insert {
					if act.Before == nil {
						l.PushFront(act.value)
					} else {
						l.InsertAfter(act.value, act.Before)
					}
				} else if act.TypeOp == Delete {
					curr := l.Front()
					for curr != nil {
						if curr != cursor && curr.Value.(rune) == act.value && curr.Prev() == act.Before {
							l.Remove(curr)
							break
						}
						curr = curr.Next()
					}
				}
			}
		}
	}

	var sb strings.Builder
	for curr := l.Front(); curr != nil; curr = curr.Next() {
		sb.WriteRune(curr.Value.(rune))
	}
	fmt.Println(sb.String())
}