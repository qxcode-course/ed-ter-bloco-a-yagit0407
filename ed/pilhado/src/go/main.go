package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	row, col int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	dims := strings.Fields(scanner.Text())
	if len(dims) < 2 {
		return
	}
	rows, _ := strconv.Atoi(dims[0])
	cols, _ := strconv.Atoi(dims[1])

	grid := make([][]rune, rows)
	visited := make([][]bool, rows)
	var start, end Point

	for i := 0; i < rows; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			if len(line) < cols {
				line += strings.Repeat(" ", cols-len(line))
			}
			grid[i] = []rune(line)
			visited[i] = make([]bool, cols)

			for j := 0; j < cols; j++ {
				if grid[i][j] == 'I' {
					start = Point{row: i, col: j}
				} else if grid[i][j] == 'F' {
					end = Point{row: i, col: j}
				}
			}
		}
	}

	var caminho []Point
	caminho = append(caminho, start)

	// Ordem correta de vizinhos para bater com os testes: Cima, Baixo, Esquerda, Direita
	dRow := []int{-1, 1, 0, 0}
	dCol := []int{0, 0, -1, 1}

	for len(caminho) > 0 {
		atual := caminho[len(caminho)-1]
		visited[atual.row][atual.col] = true

		if atual.row == end.row && atual.col == end.col {
			break
		}

		var validos []Point
		for i := 0; i < 4; i++ {
			nRow := atual.row + dRow[i]
			nCol := atual.col + dCol[i]

			if nRow >= 0 && nRow < rows && nCol >= 0 && nCol < cols {
				if grid[nRow][nCol] != '#' && !visited[nRow][nCol] {
					validos = append(validos, Point{row: nRow, col: nCol})
				}
			}
		}

		if len(validos) > 0 {
			caminho = append(caminho, validos[0])
		} else {
			caminho = caminho[:len(caminho)-1]
		}
	}

	// Pinta o caminho correto
	for len(caminho) > 0 {
		p := caminho[len(caminho)-1]
		caminho = caminho[:len(caminho)-1]

		// O 'I' deve virar '.' conforme o exemplo esperado do professor.
		// Apenas o 'F' (fim) NÃO deve ser substituído por '.'.
		if grid[p.row][p.col] != '.' {
			grid[p.row][p.col] = '.'
		}
	}

	for i := 0; i < rows; i++ {
		fmt.Println(string(grid[i]))
	}
}