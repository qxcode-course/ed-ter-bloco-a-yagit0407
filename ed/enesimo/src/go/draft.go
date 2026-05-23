package main
import (
    "fmt"
    
)

func temDivisor(n, i int) bool {
	if i*i > n {
		return false
	}
	if n%i == 0 {
		return true
	}
	return temDivisor(n, i+2)
}

func ehPrimo(n int) bool {
    if n < 2 {
        return false
    } 
    if n == 2 {
        return true
    }
     if n%2 == 0 {
        return false
    }
    return !temDivisor(n, 3)
}

func nthPrimo(n, candidato int) int {
    if ehPrimo(candidato) {
        n--
    }
    if n == 0 {
        return candidato
    }
    return nthPrimo(n, candidato+1)
}

func main() {
    var n int

    fmt.Scan(&n)
    fmt.Println(nthPrimo(n, 2))
}
