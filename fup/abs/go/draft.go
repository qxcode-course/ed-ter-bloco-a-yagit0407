package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    resultado := a - b

    if resultado < 0 {
        resultado = -resultado
    }
    fmt.Println(resultado)
}
