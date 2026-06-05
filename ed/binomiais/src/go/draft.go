package main
import "fmt"
func main() {
    var n, k int

    fmt.Scan(&n, &k)

    if k < 0 || k > n {
        fmt.Println(0)
    }
    if k == 0 || k == n {
        fmt.Println(1)
        return
    }
    if k > n-k {
        k = n - k
    }
    dp := make([]int, k + 1)
    dp[0] = 1

    for i := 1; i <= n; i++ {
        limite := i
        if limite > k {
            limite = k
        }
        for j := limite; j > 0; j-- {
            dp[j] = dp[j] + dp[j-1]
        }
    }
    fmt.Println(dp[k])
}