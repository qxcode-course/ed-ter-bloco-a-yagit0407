package main
import "fmt"
func main() {
    var a, b int
    var c string
    fmt.Scan(&a, &b, &c)
    if c == "+" {
        fmt.Println(a + b)
    } else if c == "-" {
        fmt.Println(a - b)
    } else if c == "*" {
        fmt.Println(a * b)
    } else if c == "/" {
        fmt.Println(a / b)
    }
}
