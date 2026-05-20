package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	nrows := len(board)
	ncols := len(board[0])

	
	for j := 0; j < ncols; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}
		if board[nrows-1][j] == 'O' {
			dfs(board, nrows-1, j)
		}
	}


	for i := 0; i < nrows; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}

		if board[i][ncols-1] == 'O' {
			dfs(board, i, ncols-1)
		}
	}

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, r, c int) {
	nrows := len(board)
	ncols := len(board[0])

	// Trocado '0' por 'O'
	if r < 0 || r >= nrows || c < 0 || c >= ncols || board[r][c] != 'O' {
		return
	}
	board[r][c] = '*'

	dfs(board, r-1, c)
	dfs(board, r+1, c)
	dfs(board, r, c-1)
	dfs(board, r, c+1)
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
