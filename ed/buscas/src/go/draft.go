package main

import (
	"bufio"
	"fmt"
	"os"
)

func matchingStrings(strings []string, queries []string) []int {
	frequencia := make(map[string]int)

	for _, s := range strings {
		frequencia[s]++
	}

	resultado := make([]int, 0)

	for _, q := range queries {
		resultado = append(resultado, frequencia[q])
	}

	return resultado
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	strings := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &strings[i])
	}

	var m int
	fmt.Fscan(in, &m)

	queries := make([]string, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &queries[i])
	}

	resultado := matchingStrings(strings, queries)

	for i, v := range resultado {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
    fmt.Println()
}