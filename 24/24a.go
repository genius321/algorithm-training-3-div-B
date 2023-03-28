package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func min(data ...int) int {
	sort.Ints(data)
	return data[0]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	dp := make([]int, n+3)
	abc := make([][3]int, n+3)
	for i := 0; i < 3; i++ {
		abc[i][0] = 3601
		abc[i][1] = 3601
		abc[i][2] = 3601
	}
	for i := 3; i < n+3; i++ {
		fmt.Fscan(in, &abc[i][0], &abc[i][1], &abc[i][2])
	}
	for i := 3; i < n+3; i++ {
		dp[i] = min(dp[i-1]+abc[i][0],
			dp[i-2]+abc[i-1][1],
			dp[i-3]+abc[i-2][2])
	}
	fmt.Fprintln(out, dp[n+2])
}
