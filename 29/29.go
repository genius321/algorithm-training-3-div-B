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

	if n == 0 {
		fmt.Fprintln(out, 0)
		fmt.Fprintln(out, 0, 0)
		return
	}

	data := make([]int, n)
	var coup int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &data[i])
		if data[i] > 100 {
			coup++
		}
	}

	dp := make([][]int, n+1)
	// 0 купонов
	for i := range dp {
		//учитваю случай, когда купонов 0
		dp[i] = make([]int, coup+2)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < coup+2; j++ {
			dp[i][j] = 30000
		}
	}
	for i := 1; i < n+1; i++ {
		dp[i][0] = i
	}
	for i := 2; i < coup+2; i++ {
		dp[0][i] = i - 1
	}
	var maxCoup int
	if data[0] <= 100 {
		dp[1][1] = data[0]
	} else {
		dp[1][2] = data[0]
		maxCoup++
	}
	for i, l := 1, len(data); i < l; i++ {
		for j := 0; j <= coup; j++ {
			a := 30000
			b := 30000
			if j+1 <= maxCoup {
				a = dp[i][j+2]
			}
			if j-1 >= 0 && data[i] > 100 {
				b = dp[i][j] + data[i]
			}
			if data[i] <= 100 {
				b = dp[i][j+1] + data[i]
			}
			dp[i+1][j+1] = min(a, b)
		}
		if data[i] > 100 {
			maxCoup++
		}
	}

	tmp := make([]int, len(dp[n][1:]))
	copy(tmp, dp[n][1:])

	minSum := min(tmp...)
	fmt.Fprintln(out, minSum)

	var remainingCoups int
	for i := len(dp[n]) - 1; i > 0; i-- {
		if dp[n][i] == minSum {
			remainingCoups = i - 1
			break
		}
	}
	fmt.Fprint(out, remainingCoups, " ")

	dayUsingCoup := make([]int, 0)

	for n != 1 || (remainingCoups+1 != 2 && remainingCoups+1 != 1) {
		if data[n-1] > 100 {
			if remainingCoups+2 <= coup+1 && dp[n][remainingCoups+1] == dp[n-1][remainingCoups+2] {
				dayUsingCoup = append(dayUsingCoup, n)
				n--
				remainingCoups++
			} else {
				n--
				remainingCoups--
			}
		} else {
			if remainingCoups+2 <= coup+1 && dp[n][remainingCoups+1] == dp[n-1][remainingCoups+2] {
				dayUsingCoup = append(dayUsingCoup, n)
				n--
				remainingCoups++
			} else {
				n--
			}
		}
	}
	fmt.Fprintln(out, len(dayUsingCoup))
	for i := len(dayUsingCoup) - 1; i >= 0; i-- {
		fmt.Fprintln(out, dayUsingCoup[i])
	}
}
