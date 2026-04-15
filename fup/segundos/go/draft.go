package main
import "fmt"
func main() {
    var segundos int
    fmt.Scan(&segundos)

    horas := segundos / 3600
    resto := segundos % 3600
    minutos := resto / 60
    segundos = resto % 60

    fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)
}
