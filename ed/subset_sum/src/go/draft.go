package main
import "fmt"

func subsetSum(vet []int, index int, somaAtual int, alvo int) bool {

    if somaAtual == alvo {
        return true
    }

    if index >= len(vet) {
        return false
    }

    if subsetSum(vet, index+1, somaAtual + vet[index], alvo) {
        return true
    }

    if subsetSum(vet, index + 1, somaAtual, alvo) {
        return true
    }
    return false
}

func main() {

    var n, k int

    fmt.Scan(&n, &k)

    vet := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&vet[i])
    }
    fmt.Println(subsetSum(vet, 0, 0, k))
}