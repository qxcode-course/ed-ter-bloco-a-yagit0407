package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    N := make([]int, n)  
    for i := 0; i < n; i++ {
        fmt.Scan(&N[i])
    }

    casal := 0

    for i := 0; i < n-1; {
        if N[i]*N[i+1] < 0 {
            casal++
            i += 2
        } else {
            i++
        }
    }

    fmt.Println(casal)
}