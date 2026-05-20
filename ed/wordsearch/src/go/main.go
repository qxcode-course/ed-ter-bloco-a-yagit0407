package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	
	if len(grid) == 0 {
		return false
	}

	nl := len(grid)
	nc := len(grid[0])

	var dfs func(int, int, int) bool

	dfs = func(l, c, pos int) bool {
		
		if pos == len(word) {
			return true
		}

		if l < 0 || l >= nl || c < 0 || c >= nc {
			return false
		}
		
		if grid[l][c] != word[pos] {
			return false
		}
		
		temp := grid[l][c]
		grid[l][c] = '#'

		found := 
		dfs(l+1, c, pos+1) ||
		dfs(l-1, c, pos+1) ||
		dfs(l, c+1, pos+1) ||
		dfs(l, c-1, pos+1)
			
		grid[l][c] = temp

		return found
	}

	for i := 0; i < nl; i++ {
		for j := 0; j < nc; j++ {

			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
