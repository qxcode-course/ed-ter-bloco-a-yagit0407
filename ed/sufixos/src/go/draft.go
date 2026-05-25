package main

import (
	"fmt"
)

func imprimirSufixos(s string) {
	if len(s) == 0 {
		return
	}
	imprimirSufixos(s[1:])
	fmt.Println(s)
}

func main() {
	var palavra string
	if _, err := fmt.Scan(&palavra); err == nil {
		imprimirSufixos(palavra)
	}
}