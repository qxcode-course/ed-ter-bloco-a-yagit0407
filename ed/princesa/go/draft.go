package main
import "fmt"

func imprimir(vet[] int, espada int) {
    fmt.Print("[ ")
    for i := 0 ; i < len(vet); i++ {
        if i == espada {
            fmt.Printf("%d> ", vet[i])
        } else {
            fmt.Printf("%d ", vet[i])
        }
    }
    fmt.Println("]")
}

func main() {
    var N, E int
    fmt.Scan(&N, &E)
    vet := make([]int, N)
    for i := 0; i < N; i++ {
        vet[i] = i + 1
    }
    espada := E - 1

    for len(vet) > 1 {
        imprimir(vet, espada)
        morto := (espada + 1) % len(vet)
        vet = append(vet[:morto], vet[morto+1:]...)
        if morto < espada {
            espada--
        }
        espada = (espada + 1) % len(vet)
    }
    imprimir(vet, espada)
}
