package main
import "fmt"

func resolver(n int) {
    if n == 0 {
        return
    }
    resolver (n / 2)

    fmt.Printf("%d %d\n", n/2, n%2)
}

func main() {
    var n int
    fmt.Scan(&n)

    resolver(n)
    
}
