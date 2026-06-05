package main
import "fmt"
func main() {
    var n int

    fmt.Scan(&n)

    if n == 1 || n == 2 {
        fmt.Println(1)
        return
    } 
    if n == 3 {
        fmt.Println(2)
        return
    }
    dp := make([]int, n+1)
    dp[1] = 1
    dp[2] = 1
    dp[3] = 2

    for i := 4; i <= n; i++ {
        dp[i] = dp[i - 1] + dp[i - 3]
    }
    fmt.Println(dp[n])
}