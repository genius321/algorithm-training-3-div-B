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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	data := make([][]int, n)
	for i := range data {
		data[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &data[i][j])
		}
	}

	prefixsum := make([][]int, n)
	for i := range data {
		prefixsum[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 1; j < m+1; j++ {
			prefixsum[i][j] = prefixsum[i][j-1] + data[i][j-1]
		}
	}

	for i := 0; i < k; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		var res int
		for i := x1 - 1; i < x2; i++ {
			res += prefixsum[i][y2] - prefixsum[i][y1-1]
		}
		fmt.Fprintln(out, res)
	}
}
