package main
import "fmt"
func main() {
    var a string
    var idade int

    fmt.Scan(&a, &idade)

    if idade < 12 {
        fmt.Println(a, "eh crianca")
    } else if idade < 18 {
        fmt.Println(a,"eh jovem")
    } else if idade < 65 {
        fmt.Println(a,"eh adulto")
    } else if idade < 1000 {
        fmt.Println(a, "eh idoso")
    } else {
        fmt.Println(a, "eh mumia")
    }
}
