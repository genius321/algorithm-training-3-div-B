package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func min(s ...int) int {
	sort.Ints(s)
	return s[0]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	dp := make([]int, n+3)
	dp[1] = 0
	dp[2] = 1
	dp[3] = 1
	for i, l := 4, n+1; i < l; i++ {
		var a, b int
		if i%3 == 0 {
			a = i / 3
		}
		if i%2 == 0 {
			b = i / 2
		}
		if a > 0 && b > 0 {
			dp[i] = min(dp[a], dp[b], dp[i-1]) + 1
		} else if a > 0 {
			dp[i] = min(dp[a], dp[i-1]) + 1
		} else if b > 0 {
			dp[i] = min(dp[b], dp[i-1]) + 1
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	fmt.Println(dp[n])

	res := make([]int, 0)
	res = append(res, n)
	for i := dp[n]; i > 1; i-- {
		var a, b int
		if n%3 == 0 {
			a = n / 3
		}
		if n%2 == 0 {
			b = n / 2
		}
		var m int
		if a > 0 && b > 0 {
			m = min(dp[a], dp[b], dp[n-1])
		} else if a > 0 {
			m = min(dp[a], dp[n-1])
		} else if b > 0 {
			m = min(dp[b], dp[n-1])
		} else {
			m = dp[n-1]
		}
		if m == dp[a] {
			res = append(res, a)
			n /= 3
		} else if m == dp[b] {
			res = append(res, b)
			n /= 2
		} else {
			res = append(res, n-1)
			n--
		}
	}
	if n > 1 {
		fmt.Print(1, " ")
	}
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i], " ")
	}
	fmt.Println()
}
