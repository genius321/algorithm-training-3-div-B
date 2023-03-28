package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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
	m := make(map[int]bool)
	for _, v := range data {
		m[v] = true
	}
	uniqData := make([]int, len(m))
	i := 0
	for k := range m {
		uniqData[i] = k
		i++
	}
	sort.Ints(uniqData)
	var k int
	fmt.Fscan(in, &k)
	for i := 0; i < k; i++ {
		var tmp int
		fmt.Fscan(in, &tmp)
		var m int
		if uniqData[0] >= tmp {
			fmt.Fprintln(out, 0)
			continue
		}
		l := 0
		r := len(uniqData) - 1
		for l <= r {
			m = (l + r) / 2
			if uniqData[m] < tmp {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		fmt.Fprintln(out, l)
	}
}
