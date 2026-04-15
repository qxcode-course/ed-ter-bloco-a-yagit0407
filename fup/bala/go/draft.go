package main
import "fmt"
import "math"
func main() {
    var x, y, x2, y2 float64
    fmt.Scan(&x, &y, &x2, &y2)
    d := math.Sqrt(math.Pow((x2 - x), 2) + math.Pow((y2 - y), 2))
    fmt.Printf("%.2f\n", d)

}
