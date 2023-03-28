package main

import (
	"bufio"
	"fmt"
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

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &dp[i][j])
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] += min(dp[i][j-1], dp[i-1][j])
			} else if i-1 >= 0 {
				dp[i][j] += dp[i-1][j]
			} else if j-1 >= 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	fmt.Fprintln(out, dp[n-1][m-1])
}
