package main

import "fmt"
func main() {
    var a string
    var c string
    fmt.Scan(&a)
    var b int
    fmt.Scan(&b)
    if(b < 12) {
        c = "crianca"
    } else if b < 18 {
        c = "jovem"
    } else if b < 65 {
        c = "adulto"
    } else if b < 1000 {
    c = "idoso"
    } else  {
    c = "mumia"
    }
    fmt.Printf("%s eh %s\n", a, c)
}
