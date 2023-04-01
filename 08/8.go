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

	var k int
	fmt.Fscan(in, &k)

	data := make([][2]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &data[i][0], &data[i][1])
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	x1 := data[0][0]
	x2 := data[len(data)-1][0]
	sort.Slice(data, func(i, j int) bool {
		return data[i][1] < data[j][1]
	})
	y1 := data[0][1]
	y2 := data[len(data)-1][1]
	fmt.Fprintln(out, x1, y1, x2, y2)
}
