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

	var k int
	var s string
	fmt.Fscan(in, &k, &s)

	res := 0
	for i := 0; i < 26; i++ {
		c := 0
		ch := byte('a' + i)
		var ptr1, ptr2 int
		for ; ptr1 < len(s); ptr1++ {
			for ptr2 < len(s) && c <= k {
				if s[ptr2] != ch {
					c++
				}
				if c > k {
					break
				}
				ptr2++
			}
			if res < ptr2-ptr1 {
				res = ptr2 - ptr1
			}
			if s[ptr1] != ch {
				c -= 2
			}
		}
	}
	fmt.Fprintln(out, res)
}
