package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	perimetro := func(a, b, c float64) float64 {
		return (a + b + c) / 2
	}

	area := func(s, a, b, c float64) float64 {
		return math.Sqrt(s * (s - a) * (s - b) * (s - c))
	}

	p := perimetro(a, b, c)
	fmt.Printf("%.2f\n", area(p, a, b, c))
}
