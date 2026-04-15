package main
import "fmt"
import "math"
func main() {
    var n int
    fmt.Scan(&n)
    
    vencedor := -1
    melhorpontuacao := math.MaxInt64
    
    for i := 0; i <= n; i++ {
        var a, b int
        fmt.Scan(&a, &b)

        if a >= 10 && b >= 10 {
            pontuacao := a - b
            if pontuacao < 0 {
                pontuacao = -pontuacao
            }
            if vencedor == 0 || pontuacao < melhorpontuacao {
                melhorpontuacao = pontuacao
                vencedor = i
            } 
        }
    }
    if vencedor == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(vencedor)
    }
}