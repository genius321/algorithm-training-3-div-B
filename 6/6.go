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

	var m, n int
	fmt.Fscan(in, &m, &n)

	data := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &data[i][0], &data[i][1])
	}

	res := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if data[i][1] < data[j][0] || data[j][1] < data[i][0] {
				continue
			} else {
				res--
				break
			}
		}
	}
	fmt.Fprintln(out, res)
}
