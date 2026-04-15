package main
import "fmt"
func main() {
    var a, b, c int
    var d, e, f float64
    var g float64

    fmt.Scan(&a, &b, &c)
    fmt.Scan(&d, &e, &f)
    fmt.Scan(&g)

    quantidade := (float64(a) * d) + (float64(b) * e) + (float64(c) * f)
    troco := g - quantidade

    fmt.Printf("%.2f\n", troco)
}
