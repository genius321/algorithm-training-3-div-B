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
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &data[i])
	}
	l := len(data)
	for i := l/2 - 1; i >= 0; i-- {
		pos := i
		for pos*2+1 < l {
			max_son_index := pos*2 + 1
			if pos*2+2 < l && data[pos*2+2] > data[max_son_index] {
				max_son_index = pos*2 + 2
			}
			if data[pos] < data[max_son_index] {
				data[pos], data[max_son_index] = data[max_son_index], data[pos]
				pos = max_son_index
			} else {
				break
			}
		}
	}
	for i := 0; i < l-1; i++ {
		data[l-i-1], data[0] = data[0], data[l-i-1]
		heap := data[:l-i-1]
		pos := 0
		for pos*2+1 < len(heap) {
			max_son_index := pos*2 + 1
			if pos*2+2 < len(heap) && heap[pos*2+2] > heap[max_son_index] {
				max_son_index = pos*2 + 2
			}
			if heap[pos] < heap[max_son_index] {
				heap[pos], heap[max_son_index] = heap[max_son_index], heap[pos]
				pos = max_son_index
			} else {
				break
			}
		}
	}
	for _, v := range data {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
}
