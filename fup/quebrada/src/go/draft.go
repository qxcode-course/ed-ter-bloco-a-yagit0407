package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	resto := a % b
	divInteira := a / b
	divFloat := float64(a) / float64(b)

	fmt.Println(divInteira)
	fmt.Println(resto)
	fmt.Printf("%.2f\n", divFloat)
}