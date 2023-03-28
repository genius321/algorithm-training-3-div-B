package main

import (
	"bufio"
	"fmt"
	"os"
)

var tmp = -1
var ok bool

func dfs(adjacencyMatrix [][]int, visited []int, cycle *[]int, now, prev int) {
	visited[now] = 1
	for i, v := range adjacencyMatrix[now] {
		if v == 1 {
			if visited[i] == 0 {
				dfs(adjacencyMatrix, visited, cycle, i, now)
			} else if i != -1 && i != prev && visited[i] == 1 {
				tmp = i
				ok = true
			}
			if ok {
				break
			}
		}
	}
	visited[now] = 2
	if tmp != -1 {
		*cycle = append(*cycle, now)
	}
	if tmp == now {
		tmp = -1
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
	visited := make([]int, n)
	cycle := make([]int, 0)
	for i := 0; i < n; i++ {
		if ok {
			break
		}
		if visited[i] == 0 {
			dfs(adjacencyMatrix, visited, &cycle, i, -1)
		}
	}
	if ok {
		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, len(cycle))
		for _, v := range cycle {
			fmt.Fprint(out, v+1, " ")
		}
		fmt.Fprintln(out)
	} else {
		fmt.Fprintln(out, "NO")
	}
}
