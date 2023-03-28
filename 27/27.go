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

func reverse(s []rune) []rune {
	c := len(s) / 2
	for i, l := 0, len(s)-1; i < c; i++ {
		s[i], s[l-i] = s[l-i], s[i]
	}
	return s
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
	res := make([]rune, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &dp[i][j])
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] += max(dp[i][j-1], dp[i-1][j])
			} else if i-1 >= 0 {
				dp[i][j] += dp[i-1][j]
			} else if j-1 >= 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	for i, j := n-1, m-1; i > 0 || j > 0; {
		if j-1 >= 0 && i-1 >= 0 {
			if dp[i][j-1] > dp[i-1][j] {
				res = append(res, 'R')
				j--
			} else {
				res = append(res, 'D')
				i--
			}
		} else if i-1 >= 0 {
			res = append(res, 'D')
			i--
		} else if j-1 >= 0 {
			res = append(res, 'R')
			j--
		}
	}

	fmt.Fprintln(out, dp[n-1][m-1])
	res = reverse(res)
	for _, v := range res {
		fmt.Fprintf(out, "%c ", v)
	}
	fmt.Fprintln(out)
}
