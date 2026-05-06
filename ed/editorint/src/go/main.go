package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	texto   *List[*List[rune]]
	it_line *Node[*List[rune]]
	it_char *Node[rune]
	screen  tcell.Screen
	style   tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.it_char = e.it_line.Value.Insert(e.it_char, r)
	e.it_char = e.it_char.Next()
}

func (e *Editor) ColunaAtual() int {
	col := 0

	for char := e.it_line.Value.Front(); char != e.it_char; char = char.Next() {
		col++
	}

	return col
}

func (e *Editor) IrParaColuna(col int) {
	e.it_char = e.it_line.Value.Front()

	for i := 0; i < col && e.it_char != e.it_line.Value.End(); i++ {
		e.it_char = e.it_char.Next()
	}
}

func (e *Editor) KeyLeft() {
	if e.it_char.Prev() != e.it_line.Value.End() {
		e.it_char = e.it_char.Prev()
	} else if e.it_line.Prev() != e.texto.End() {
		e.it_line = e.it_line.Prev()
		e.it_char = e.it_line.Value.End()
	}
}

func (e *Editor) KeyRight() {
	if e.it_char != e.it_line.Value.End() {
		e.it_char = e.it_char.Next()
	} else if e.it_line.Next() != e.texto.End() {
		e.it_line = e.it_line.Next()
		e.it_char = e.it_line.Value.Front()
	}
}

func (e *Editor) KeyUp() {
	if e.it_line.Prev() != e.texto.End() {
		col := e.ColunaAtual()

		e.it_line = e.it_line.Prev()
		e.IrParaColuna(col)
	}
}

func (e *Editor) KeyDown() {
	if e.it_line.Next() != e.texto.End() {
		col := e.ColunaAtual()

		e.it_line = e.it_line.Next()
		e.IrParaColuna(col)
	}
}

func (e *Editor) KeyEnter() {
	nova := NewList[rune]()

	for e.it_char != e.it_line.Value.End() {
		atual := e.it_char
		e.it_char = e.it_char.Next()

		nova.Insert(nova.End(), atual.Value)
		e.it_line.Value.Erase(atual)
	}

	e.texto.Insert(e.it_line.Next(), nova)

	e.it_line = e.it_line.Next()
	e.it_char = e.it_line.Value.Front()
}

func (e *Editor) KeyBackspace() {
	if e.it_char.Prev() != e.it_line.Value.End() {
		apagar := e.it_char.Prev()
		e.it_line.Value.Erase(apagar)
	} else if e.it_line.Prev() != e.texto.End() {
		linhaAtual := e.it_line
		linhaAnterior := e.it_line.Prev()

		e.it_line = linhaAnterior
		e.it_char = linhaAnterior.Value.End()

		for char := linhaAtual.Value.Front(); char != linhaAtual.Value.End(); char = char.Next() {
			linhaAnterior.Value.Insert(linhaAnterior.Value.End(), char.Value)
		}

		e.texto.Erase(linhaAtual)
	}
}

func (e *Editor) KeyDelete() {
	if e.it_char != e.it_line.Value.End() {
		apagar := e.it_char
		e.it_char = e.it_char.Next()
		e.it_line.Value.Erase(apagar)
	} else if e.it_line.Next() != e.texto.End() {
		proximaLinha := e.it_line.Next()

		for char := proximaLinha.Value.Front(); char != proximaLinha.Value.End(); char = char.Next() {
			e.it_line.Value.Insert(e.it_line.Value.End(), char.Value)
		}

		e.texto.Erase(proximaLinha)
	}
}

func main() {
	editor := NewEditor()
	defer editor.screen.Fini()

	editor.Draw()
	editor.MainLoop()
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return

			case tcell.KeyEnter:
				e.KeyEnter()

			case tcell.KeyLeft:
				e.KeyLeft()

			case tcell.KeyRight:
				e.KeyRight()

			case tcell.KeyUp:
				e.KeyUp()

			case tcell.KeyDown:
				e.KeyDown()

			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()

			case tcell.KeyDelete:
				e.KeyDelete()

			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}

			e.Draw()

		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}

	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}

	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}

	e.screen = screen

	e.texto = NewList[*List[rune]]()
	e.texto.PushBack(NewList[rune]())

	e.it_line = e.texto.Front()
	e.it_char = e.it_line.Value.End()

	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	e.screen.SetStyle(e.style)
	e.screen.Clear()

	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()

	x := 0
	y := 0

	for line := e.texto.Front(); line != e.texto.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value

			if char == line.Value.End() {
				data = '⤶'
			}

			if data == ' ' {
				data = '·'
			}

			if char == e.it_char {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}

			x++

			if char == line.Value.End() {
				break
			}
		}

		y++
		x = 0
	}

	e.screen.Show()
}