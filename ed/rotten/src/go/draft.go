package main
import "fmt"

type Node struct {
    l, c int
    minuto int
}

func orangeRotting(grid[][] int) int {
    linhas := len(grid)
    colunas := len(grid)

    queue := []Node{}
    laranjasFrescas := 0

    for l := 0; l < linhas; l++ {
        for c := 0; c < colunas; c++ {
            if grid[l][c] == 2 {
                queue = append(queue, Node{l, c, 0})
            } else if grid[l][c] == 1 {
                laranjasFrescas++
            }
        }
    }
    if laranjasFrescas == 0 {
        return 0
    }
    return 0

    direcoes := [][]int {{-1, 0}, {1, 0}, {0, - 1}, {0, 1}}
    minutos := 0

    for len(queue) > 0 {

    }
}



func main() {
    var m, n int
    fmt.Scan(&m, &n)

    grid := make([][]int, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&grid[i][j])
        }
    }
    resultado := orangeRotting(grid)
    fmt.Println(resultado)
}