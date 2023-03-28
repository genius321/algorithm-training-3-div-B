package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, c string
	fmt.Fscan(in, &a, &b, &c)

	layout := time.TimeOnly
	a1, _ := time.Parse(layout, a)
	b1, _ := time.Parse(layout, b)
	c1, _ := time.Parse(layout, c)
	ok := a1.Compare(c1)
	diff := c1.Sub(a1)
	if ok == 1 {
		diff -= 24 * time.Hour
	}
	diff = diff / 2
	res := b1.Add(diff).Round(time.Second)
	fmt.Fprintln(out, res.Format(time.TimeOnly))
}
