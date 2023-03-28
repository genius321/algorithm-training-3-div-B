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

	var n int
	fmt.Fscan(in, &n)

	var res, prev int
	for i := 0; i < n; i++ {
		var now int
		fmt.Fscan(in, &now)
		res += min(prev, now)
		prev = now
	}

	fmt.Fprintln(out, res)
}
