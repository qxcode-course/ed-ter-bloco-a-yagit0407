package main
import "fmt"
func main() {
    var q int
    var d string
    fmt.Scan(&q, &d)

    x:= make([]int, q)
    y:= make([]int, q)
    for i := 0; i < q; i++ {
        fmt.Scan(&x[i], &y[i])
    }
    olX := make([]int, q)
    olY := make([]int, q)
    copy(olX, x)
    copy(olY, y)
    switch d { 
        case "L":
            x[0] --
        case "R":
            x[0] ++
            case "U":
            y[0] --
        case "D":
            y[0] ++
    }
    for i := 1 ; i < q; i++ {
        x[i] = olX[i-1]
        y[i] = olY[i-1]
    }
    for i := 0; i < q; i++ {
        fmt.Printf("%d %d\n", x[i], y[i])
    }
}
