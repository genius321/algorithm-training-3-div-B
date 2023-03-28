package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []int

func (ptr *stack) push(n int) {
	*ptr = append(*ptr, n)
}

func (ptr *stack) pop() int {
	n := (*ptr)[len(*ptr)-1]
	*ptr = (*ptr)[:len(*ptr)-1]
	return n
}

func (ptr *stack) back() int {
	n := (*ptr)[len(*ptr)-1]
	return n
}

func (ptr *stack) size() int {
	n := len(*ptr)
	return n
}

func (ptr *stack) clear() {
	*ptr = nil
}

func (ptr *stack) exit() string {
	return "bye"
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	st := stack{}
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
		case "back":
			if len(st) > 0 {
				n = st.back()
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
