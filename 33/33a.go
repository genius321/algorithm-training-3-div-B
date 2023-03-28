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

	var n, m int
	fmt.Fscan(in, &n, &m)

	m1 := make(map[int]bool)
	m2 := make(map[int]bool)

	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Fscan(in, &v1, &v2)
		m1[v1] = true
		m2[v2] = true
	}

	for k := range m1 {
		if m2[k] {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
}
