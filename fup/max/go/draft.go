package main
import "fmt"
func main() {
    var n, m int
    fmt.Scan(&n, &m)
    
    if n > m {
        fmt.Printf("%d\n", n)
    } else {
        fmt.Printf("%d\n", m)
    }
    
}
