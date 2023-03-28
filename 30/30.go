package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var m int
	fmt.Fscan(in, &m)

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		dp[i][0] = a[i-1]
	}
	for i := 1; i < m+1; i++ {
		dp[0][i] = b[i-1]
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if dp[i][0] == dp[0][j] {
				dp[i][j] = 1
				if i-1 >= 1 && j-1 >= 1 {
					dp[i][j] += dp[i-1][j-1]
				}
			} else {
				if i-1 >= 1 && j-1 >= 1 {
					dp[i][j] += max(dp[i-1][j], dp[i][j-1])
				} else if i-1 >= 1 {
					dp[i][j] += dp[i-1][j]
				} else if j-1 >= 1 {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}

	res := make([]int, 0)
	for i := dp[n][m]; i > 0; i-- {
		for {
			if n-1 >= 1 && dp[n][m] == dp[n-1][m] {
				n--
			} else if m-1 >= 1 && dp[n][m] == dp[n][m-1] {
				m--
			} else {
				break
			}
		}
		res = append(res, dp[0][m])
		n--
		m--
	}

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Fprint(out, res[i], " ")
	}
	fmt.Fprintln(out)
}
