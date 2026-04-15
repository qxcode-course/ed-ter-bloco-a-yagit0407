package main

import "fmt"

func main() {
	var total, qtd int
	var v [100]int
	var tem [51]int

	fmt.Scan(&total)
	fmt.Scan(&qtd)

	for i := 0; i < qtd; i++ {
		fmt.Scan(&v[i])
		tem[v[i]]++
	}

	achou := false

	for i := 1; i < qtd; i++ {
		if v[i] == v[i-1] {
			if achou {
				fmt.Print(" ")
			}
			fmt.Print(v[i])
			achou = true
		}
	}

	if !achou {
		fmt.Print("N")
	}

	fmt.Println()

	achou = false

	for i := 1; i <= total; i++ {
		if tem[i] == 0 {
			if achou {
				fmt.Print(" ")
			}
			fmt.Print(i)
			achou = true
		}
	}

	if !achou {
		fmt.Print("N")
	}

	fmt.Println()
}
