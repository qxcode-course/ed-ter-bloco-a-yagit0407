package main
import (
    "fmt"
    "strings"
)

func diagonal(s string, k int) {
    if len(s) == 0 {
        return
    }

    fmt.Print(strings.Repeat(" ", k))
    fmt.Printf("%c\n", s[0])
    diagonal(s[1:], k+1)
}

func main() {
    var palavra string
    fmt.Scan(&palavra)
    diagonal(palavra, 0)
}
