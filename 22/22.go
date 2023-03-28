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

	var n, k int
	fmt.Fscan(in, &n, &k)

	stairs := make([]int, n)

	//не прыгать - тоже способ
	stairs[0] = 1
	for i := 1; i < n; i++ {
		for j := 1; j <= k && i-j >= 0; j++ {
			stairs[i] += stairs[i-j]
		}
	}

	fmt.Fprintln(out, stairs[n-1])
}
