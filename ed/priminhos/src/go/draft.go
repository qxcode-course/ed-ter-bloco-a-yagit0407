package main
import "fmt"

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

func gerarPrimos(n, candidato int, vetor []int) []int {
    if len(vetor) == n {
        return vetor
    }
    if ehPrimo(candidato) {
        vetor = append(vetor, candidato)
    }
    return gerarPrimos(n, candidato+1, vetor)
}

func printSlice(vet []int) {
    fmt.Print("[")
    for i, x := range vet {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Print(x)
    }
    fmt.Println("]")

}
func main() {
    var n int
    if _, err := fmt.Scan(&n); err == nil {
        if n <= 0 {
            fmt.Println("[]")
            return 
        }
        primos := gerarPrimos(n, 2, []int{})
        printSlice(primos)
    }
}
