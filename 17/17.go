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
	q1 := make([]int, 5)
	q2 := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscan(in, &q1[i])
	}
	for i := 0; i < 5; i++ {
		fmt.Fscan(in, &q2[i])
	}
	for i := 0; i < 1_000_000; i++ {
		if q1[0] == 0 && q2[0] == 9 {
			q1 = append(q1[1:], q1[0])
			q1 = append(q1, q2[0])
			q2 = q2[1:]
		} else if q1[0] == 9 && q2[0] == 0 {
			q2 = append(q2, q1[0])
			q2 = append(q2[1:], q2[0])
			q1 = q1[1:]
		} else if q1[0] > q2[0] {
			q1 = append(q1[1:], q1[0])
			q1 = append(q1, q2[0])
			q2 = q2[1:]
		} else {
			q2 = append(q2, q1[0])
			q2 = append(q2[1:], q2[0])
			q1 = q1[1:]
		}
		if len(q1) == 0 {
			fmt.Fprintln(out, "second", i+1)
			return
		} else if len(q2) == 0 {
			fmt.Fprintln(out, "first", i+1)
			return
		}
	}
	fmt.Fprintln(out, "botva")
}
