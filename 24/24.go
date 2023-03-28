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

	//минимальное коль-во времени требуемое для обслуживания iого покупателя
	dp := make([]int, n)
	abc := make([][3]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &abc[i][0], &abc[i][1], &abc[i][2])
	}
	dp[0] = abc[0][0]
	if n > 1 {
		dp[1] = min(abc[0][1], abc[0][0]+abc[1][0])
	}
	if n > 2 {
		dp[2] = min(abc[0][0]+abc[1][0]+abc[2][0],
			abc[0][0]+abc[1][1],
			abc[0][1]+abc[2][0],
			abc[0][2])
	}
	for i := 3; i < n; i++ {
		dp[i] = min(dp[i-1]+abc[i][0],
			dp[i-2]+abc[i-1][1],
			dp[i-3]+abc[i-2][2])
	}
	fmt.Fprintln(out, dp[n-1])
}
