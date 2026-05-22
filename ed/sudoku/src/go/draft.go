package main
import (
	"fmt"
	"math"
)

func eValido(matriz [][]rune, lin, col int, num rune) bool {
    dim := len(matriz)

    for i := 0; i < dim; i++ {
        if matriz[lin][i] == num || matriz[i][col] == num {
            return false
        }
    }

    regiaoSize := int(math.Sqrt(float64(dim)))
    startLin := (lin / regiaoSize) * regiaoSize
    startCol := (col / regiaoSize) * regiaoSize

    for i := 0; i < regiaoSize; i++ {
        for j := 0; j < regiaoSize; j++ {
            if matriz[startLin + i][startCol + j] == num {
                return false
            }
        }
    }
    return true
}

func resolver(matriz [][]rune, index int) bool {
    dim := len(matriz)

    if index == dim*dim {
        return true
    }

    l := index / dim
    c := index % dim

    if matriz[l][c] != '.' {
        return resolver(matriz, index + 1)
    }

    for num := '1'; num < '1' + rune (dim); num++ {
        if eValido(matriz, l, c, num) {
            matriz[l][c] = num

            if resolver(matriz, index + 1) {
                return true
            }
            matriz[l][c] = '.'
        }
    }
    return false
}
func main() {
    var n int

    if _, err := fmt.Scan(&n); err != nil {
        return
    }

    matriz := make([][]rune, n)

    for i := 0; i < n; i++ {
        var linha string
        fmt.Scan(&linha)
        matriz[i] = []rune(linha)
    }
    if resolver(matriz, 0) {
        for i := 0; i < n; i++ {
            fmt.Println(string(matriz[i]))
        }
    }
}
