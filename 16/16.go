package main

import (
	"bufio"
	"fmt"
	"os"
)

type queue []int

func (ptr *queue) push(n int) {
	*ptr = append(*ptr, n)
}

func (ptr *queue) pop() int {
	n := (*ptr)[0]
	*ptr = (*ptr)[1:]
	return n
}

func (ptr *queue) front() int {
	n := (*ptr)[0]
	return n
}

func (ptr *queue) size() int {
	n := len(*ptr)
	return n
}

func (ptr *queue) clear() {
	*ptr = nil
}

func (ptr *queue) exit() string {
	return "bye"
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	st := queue{}
	for {
		_, err := fmt.Fscan(in, &s)
		if err != nil {
			break
		}
		var n int
		switch s {
		case "push":
			fmt.Fscan(in, &n)
			st.push(n)
			fmt.Fprintln(out, "ok")
		case "pop":
			if len(st) > 0 {
				n = st.pop()
				fmt.Fprintln(out, n)
			} else {
				fmt.Fprintln(out, "error")
			}
		case "front":
			if len(st) > 0 {
				n = st.front()
				fmt.Fprintln(out, n)
			} else {
				fmt.Fprintln(out, "error")
			}
		case "size":
			n = st.size()
			fmt.Fprintln(out, n)
		case "clear":
			st.clear()
			fmt.Fprintln(out, "ok")
		case "exit":
			s = st.exit()
			fmt.Fprintln(out, s)
			return
		}
	}
}
