package main
import "fmt"

func blocos(n int) int {
    if n == 1 {
        return 20
    }
    return blocos(n - 1) + 8
}
func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(blocos(n))
}