package main
import "fmt"

type Gomo struct {
    x, y int
}
func main() {
    var q int
    var d string

    fmt.Scan(&q, &d)

    cobra := make([]Gomo, q)

    for i := 0; i < q; i++ {
        fmt.Scan(&cobra[i].x, &cobra[i].y)
    }

    for i := q - 1; i > 0; i-- {
        cobra[i] = cobra[i - 1]
    }

    switch d {
	case "L":
		cobra[0].x--
	case "R":
		cobra[0].x++
	case "U":
		cobra[0].y--
	case "D":
		cobra[0].y++
	}

    for _, gomo := range cobra {
		fmt.Printf("%d %d\n", gomo.x, gomo.y)
	}
}