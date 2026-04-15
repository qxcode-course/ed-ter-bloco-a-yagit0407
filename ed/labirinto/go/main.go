package main

import (
	"bufio"
	"fmt"
	"os"
)
type Pos struct {
    l, c int
}
func search(grid [][]rune, atual Pos, endPos Pos) bool {
	if atual.l < 0 || atual.l >= len(grid) || atual.c < 0 || atual.c >= len(grid[0]) || grid[atual.l][atual.c] != ' ' {
        return false
    }
	if atual == endPos {
        grid[atual.l][atual.c] = '.'
        return true
    }
	grid[atual.l][atual.c] = '.'
	direcoes := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for _, d := range direcoes {
        proximo := Pos{atual.l + d.l, atual.c + d.c}
        if search(grid, proximo, endPos) {
            return true
        }
    }
	grid[atual.l][atual.c] = ' '
    return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
