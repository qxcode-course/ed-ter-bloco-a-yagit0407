package main

import "fmt"

func main() {
	var a, b string

	fmt.Scan(&a, &b)

	if a == "R" && b == "P" || a == "P" && b == "S" || a == "S" && b == "R" {
		fmt.Println("jog2")
	} else if a == b {
		fmt.Println("empate")
	} else {
		fmt.Println("jog1")
	}
}
