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

	var s string
	fmt.Fscan(in, &s)

	m := make(map[byte]int)
	for i, l := 0, len(s); i < l; i++ {
		m[s[i]] += (i + 1) * (l - i)
	}
	b := make([]byte, 0)
	for k := range m {
		b = append(b, k)
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for _, v := range b {
		fmt.Fprintf(out, "%c: %d\n", v, m[v])
	}
}
