package main
import "fmt"

func coelhos(n, k int) int {
    if n == 1 || n == 2 {
        return 1
    }
    fMenos2 := 1
    fMenos1 := 1
    fAtual := 0

    for i := 3; i <= n; i++ {
        fAtual = fMenos1 + (k * fMenos2)

        fMenos2 = fMenos1
        fMenos1 = fAtual
    }
    return fAtual
}
func main() {
    var n, k  int

    fmt.Scan(&n, &k)
    fmt.Println(coelhos(n, k))
}