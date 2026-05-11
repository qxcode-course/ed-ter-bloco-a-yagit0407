package main
import (
    "fmt"
    "bufio"
    "os"
)
type Node struct {
    value rune
    prev *Node
    next *Node
}

type Editor struct {
    head *Node
    tail *Node
    cursor *Node
}

func NewEditor() *Editor {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &Editor{
		head:   head,
		tail:   tail,
		cursor: tail,
	}
}

func (e *Editor) Insert (ch rune) {
    node := &Node{value: ch}

    prev := e.cursor.prev
    
    prev.next = node
    node.prev = prev

    node.next = e.cursor
    e.cursor.prev = node

}
func (e *Editor) Left() {
	if e.cursor.prev != e.head {
		e.cursor = e.cursor.prev
	}
}

func (e *Editor) Right() {
	if e.cursor != e.tail {
		e.cursor = e.cursor.next
	}
}

func (e *Editor) Delete() {
    if e.cursor == e.tail {
        return
    }
    target := e.cursor

    target.prev.next = target.next
    target.next.prev = target.prev
    
    e.cursor = target.next
}

func (e *Editor) Backspace() {
    if e.cursor.prev == e.head {
        return
    }
    
    target := e.cursor.prev
    target.prev.next = e.cursor
    e.cursor.prev = target.prev
}

func (e *Editor) String() string {
    out := ""

    for node := e.head.next; node != e.tail; node = node.next {
        if node == e.cursor {
            out += "|"
        }
        out += string(node.value)
    } 
    if e.cursor == e.tail {
        out += "|"
    }
    return out
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text := scanner.Text()

	ed := NewEditor()

	for _, ch := range text {

		switch ch {

		case '<':
			ed.Left()

		case '>':
			ed.Right()

		case 'B':
			ed.Backspace()

		case 'D':
			ed.Delete()

		case 'R':
			ed.Insert('\n')

		default:
			ed.Insert(ch)
		}
	}

	fmt.Println(ed.String())
}
