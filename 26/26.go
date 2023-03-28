package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < 1; i++ {
		for j := 2; j <= m; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	for i := 2; i <= n; i++ {
		for j := 0; j < 1; j++ {
			dp[i][j] = math.MaxInt
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &dp[i][j])
			dp[i][j] += min(dp[i][j-1], dp[i-1][j])
		}
	}

	fmt.Fprintln(out, dp[n][m])
}
