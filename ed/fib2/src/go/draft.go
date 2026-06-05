package main
import "fmt"

type Populacao struct {
    idade0 int
    idade1 int
    idade2 int
}
var memo map[int] Populacao 

func simular(n int) Populacao {
    if n == 1 {
        return Populacao{idade0: 1, idade1: 0, idade2: 0}
    }
    if res, existe := memo[n]; existe {
        return res
    }    
    anterior := simular(n - 1)
    var atual Populacao
    atual.idade0 = anterior.idade1 + anterior.idade2
    atual.idade1 = anterior.idade0
    atual.idade2 = anterior.idade1

    memo[n] = atual
    return atual
}

func main() {
    var n int

    fmt.Scan(&n)

    memo = make(map[int]Populacao)
    popFinal := simular(n)
    total := popFinal.idade0 + popFinal.idade1 + popFinal.idade2
    
    fmt.Println(total)
}