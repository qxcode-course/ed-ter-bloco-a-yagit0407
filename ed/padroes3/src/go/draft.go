package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	resultado := ((n-2)*m*m - (n-4)*m) / 2

	fmt.Println(resultado)
}