package main

import "fmt"

func soma(a, b int) {
	fmt.Println(a + b)
}

func subtracao(a, b int) {
	fmt.Println(a - b)
}

func multiplicacao(a, b int) {
	fmt.Println(a * b)
}


func divisao(a, b int) {
	fmt.Printf("%.2f\n", float64(a)/float64(b))
}

func resto(a, b int) {
	fmt.Println(a % b)
}

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	soma(a, b)
	subtracao(a, b)
	multiplicacao(a, b)
	divisao(a, b)

	resto(a, b)
}