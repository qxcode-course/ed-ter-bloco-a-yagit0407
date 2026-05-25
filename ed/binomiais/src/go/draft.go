package main
import "fmt"

func calcularCombinacao(n, k int, memo [][]int) int {
    if k == 0 || k == n {
        return 1
    } 
    if memo[n][k] != 0 {
        return memo[n][k]
    }
    memo[n][k] = calcularCombinacao(n-1, k-1, memo) + calcularCombinacao(n-1, k, memo)
    return memo[n][k]
}
func main() {
    var n, k int
    fmt.Scan(&n, &k)

    memo := make([][]int, n+1)

    for i := range memo {
			memo[i] = make([]int, k+1)
		}
		fmt.Println(calcularCombinacao(n, k, memo))
}
