package main
import "fmt"
func main() {
    var s string
    fmt.Scan(&s)

    pilha := []rune{}
    balanceado := true

    for _, c := range s {
        if c == '(' || c == '[' {
            pilha = append(pilha, c)
        } else {
            if len(pilha) == 0 {
                balanceado = false
                break
            }

            topo := pilha[len(pilha)-1]

            if (c == ')' && topo == '(') || (c == ']' && topo == '[') {
                pilha = pilha[:len(pilha)-1]
            } else {
                balanceado = false
                break
            }

        }
    }
    if len(pilha) != 0 {
        balanceado = false
    }

    if balanceado {
        fmt.Println("balanceado")
    } else {
        fmt.Println("nao balanceado")
    }

}