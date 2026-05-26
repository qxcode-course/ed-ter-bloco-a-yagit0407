package main
import "fmt"
func main() {
    
    var t, r int

    fmt.Scan(&t, &r)

    v := make([]int, t)
    res := make([]int, t)

    for i := 0; i < t; i++ {
        fmt.Scan(&v[i])
    }

    r = r % t

    for i := 0; i < t; i++ {
        novaPos := (i + r) % t
        res[novaPos] = v[i]
    }

    fmt.Print("[ ")

    for i := 0; i < t; i++ {
		fmt.Print(res[i], " ")
	}

	fmt.Println("]")
}