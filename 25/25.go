package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var n int
	fmt.Fscan(in, &n)

	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &data[i])
	}
	sort.Ints(data)
	dp := make([]int, n)
	dp[0] = data[1] - data[0]
	dp[1] = dp[0]
	for i, l := 2, len(data); i < l; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + data[i] - data[i-1]
	}
	fmt.Fprintln(out, dp[n-1])
}
