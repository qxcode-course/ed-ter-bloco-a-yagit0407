package main
import "fmt"

type Node struct {
    l, c int
    minuto int
}

func orangeRotting(grid[][] int) int {
    linhas := len(grid)
    colunas := len(grid[0])

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
    
    

    direcoes := [][]int {{-1, 0}, {1, 0}, {0, - 1}, {0, 1}}
    minutosPassados := 0

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
       
        minutosPassados = curr.minuto
        for _, dir := range direcoes {
            nl, nc := curr.l+dir[0], curr.c+dir[1]

                if nl >= 0 && nl < linhas && nc >= 0 && nc < colunas && grid[nl][nc] == 1 {
                grid[nl][nc] = 2
                laranjasFrescas--

                queue = append(queue, Node{nl, nc, minutosPassados + 1})
            }
        }

    }
    if laranjasFrescas > 0 {
        return -1
    }

    return minutosPassados

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