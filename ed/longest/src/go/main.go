package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
	}

	var dfs func(int, int) int

	dfs = func(r, c int) int {
		if memo[r][c] != 0 {
			return memo[r][c]
		}

		best := 1

		dirs := [][]int{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		}

		for _, d := range dirs {
			nr := r + d[0]
			nc := c + d[1]

			if nr >= 0 && nr < rows &&
				nc >= 0 && nc < cols &&
				matrix[nr][nc] > matrix[r][c] {

				path := 1 + dfs(nr, nc)

				if path > best {
					best = path
				}
			}
		}

		memo[r][c] = best
		return best
	}

	ans := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cur := dfs(i, j)
			if cur > ans {
				ans = cur
			}
		}
	}

	return ans
}

				

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
