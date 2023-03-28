package main

import (
	"bufio"
	"fmt"
	"os"
)

type queue []int

func (ptr *queue) pushFront(n int) {
	*ptr = append([]int{n}, *ptr...)
}

func (ptr *queue) pushBack(n int) {
	*ptr = append(*ptr, n)
}

func (ptr *queue) popFront() int {
	n := (*ptr)[0]
	*ptr = (*ptr)[1:]
	return n
}

func (ptr *queue) popBack() int {
	n := (*ptr)[len(*ptr)-1]
	*ptr = (*ptr)[:len(*ptr)-1]
	return n
}

func (ptr *queue) front() int {
	n := (*ptr)[0]
	return n
}

func (ptr *queue) back() int {
	n := (*ptr)[len(*ptr)-1]
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
		case "push_front":
			fmt.Fscan(in, &n)
			st.pushFront(n)
			fmt.Fprintln(out, "ok")
		case "push_back":
			fmt.Fscan(in, &n)
			st.pushBack(n)
			fmt.Fprintln(out, "ok")
		case "pop_front":
			if len(st) > 0 {
				n = st.popFront()
				fmt.Fprintln(out, n)
			} else {
				fmt.Fprintln(out, "error")
			}
		case "pop_back":
			if len(st) > 0 {
				n = st.popBack()
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
