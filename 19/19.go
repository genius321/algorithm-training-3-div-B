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

	heap := make([]int, 0)
	for i := 0; i < n; i++ {
		var op int
		fmt.Fscan(in, &op)
		if op == 0 {
			var n int
			fmt.Fscan(in, &n)
			heap = append(heap, n)
			pos := len(heap) - 1
			for pos > 0 && heap[pos] > heap[(pos-1)/2] {
				heap[pos], heap[(pos-1)/2] = heap[(pos-1)/2], heap[pos]
				pos = (pos - 1) / 2
			}
		} else if op == 1 {
			fmt.Fprintln(out, heap[0])
			heap[0] = heap[len(heap)-1]
			pos := 0
			for pos*2+2 < len(heap) {
				max_son_index := pos*2 + 1
				if heap[pos*2+2] > heap[max_son_index] {
					max_son_index = pos*2 + 2
				}
				if heap[pos] < heap[max_son_index] {
					heap[pos], heap[max_son_index] = heap[max_son_index], heap[pos]
					pos = max_son_index
				} else {
					break
				}
			}
			heap = heap[:len(heap)-1]
		}
	}
}
