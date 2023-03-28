package main

import (
	"bufio"
	"fmt"
	"os"
)

var queue = make([]int, 0)
var res int
var start, end int

func bfs(adjacencyMatrix [][]int, visited []int, now int) {
	visited[now] = 0
	queue = append(queue, now)
	for len(queue) > 0 {
		for i, v := range adjacencyMatrix[queue[0]] {
			if v == 1 {
				if visited[i] == -1 {
					visited[i] = visited[queue[0]] + 1
					queue = append(queue, i)
				}
				if i == end {
					return
				}
			}
		}
		queue = queue[1:]
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	adjacencyMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		adjacencyMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &adjacencyMatrix[i][j])
		}
	}

	fmt.Fscan(in, &start, &end)
	//т.к. у меня счёт с 0
	start = start - 1
	end = end - 1

	visited := make([]int, n)
	for i := 0; i < n; i++ {
		visited[i] = -1
	}
	bfs(adjacencyMatrix, visited, start)
	fmt.Fprintln(out, visited[end])
}
