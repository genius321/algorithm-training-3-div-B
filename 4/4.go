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

	var n, k, r, p, a, b int
	fmt.Fscan(in, &n, &k, &r, &p)

	a = r * 2
	if p == 1 {
		a -= 1
	}
	b = a + k
	if b > n {
		b = a - k
	}
	if p == 2 {
		if k%2 == 1 {
			b = a - k
			if b < 1 {
				b = a + k
			}
		}
	}
	if b <= 0 {
		fmt.Fprintln(out, -1)
	} else {
		var p2 int
		if b%2 == 0 {
			p2 = 2
		} else {
			p2 = 1
		}
		var r2 int
		if p2 == 2 {
			r2 = b / 2
		} else {
			r2 = b/2 + 1
		}
		fmt.Fprintln(out, r2, p2)
	}
}
