package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, m)
	}
	// делаю пометку, что с этой клетки можно ходить
	dp[0][0][1] = 1
	// 0 способов, это и есть 1 способ, ничего не делать
	dp[0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i-2 >= 0 && j-1 >= 0 && dp[i-2][j-1][1] == 1 {
				dp[i][j][0] += dp[i-2][j-1][0]
				dp[i][j][1] = 1
			}
			if i-1 >= 0 && j-2 >= 0 && dp[i-1][j-2][1] == 1 {
				dp[i][j][0] += dp[i-1][j-2][0]
				dp[i][j][1] = 1
			}
		}
	}
	fmt.Fprintln(out, dp[n-1][m-1][0])
}
