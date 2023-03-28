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

	var n int
	fmt.Fscan(in, &n)

	path1 := make([]int, n)
	path2 := []int{}
	deadlock := []int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &path1[i])
	}
	for {
		for len(path1) > 0 {
			if len(deadlock) == 0 || deadlock[len(deadlock)-1] >= path1[0] {
				deadlock = append(deadlock, path1[0])
				path1 = path1[1:]
			} else {
				break
			}
		}
		if len(path2) == 0 ||
			(len(deadlock) > 0 && path2[len(path2)-1] <= deadlock[len(deadlock)-1]) {
			path2 = append(path2, deadlock[len(deadlock)-1])
			deadlock = deadlock[:len(deadlock)-1]
		} else if len(path1) == 0 || len(deadlock) > 0 {
			break
		}
	}
	if len(path1) == 0 && len(deadlock) == 0 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
