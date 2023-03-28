package main

import (
	"bufio"
	"fmt"
	"os"
)

func bfs(intersection [][]int, visited, start []int) {
	queue := make([]int, 0)
	for _, v := range start {
		visited[v] = 0
		queue = append(queue, v)
	}
	for len(queue) > 0 {
		for _, v := range intersection[queue[0]] {
			if visited[v] == -1 {
				visited[v] = visited[queue[0]] + 1
				queue = append(queue, v)
			}
		}
		queue = queue[1:]
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	stations := make([]map[int]bool, m)
	for i := 0; i < m; i++ {
		stations[i] = make(map[int]bool)
		var cntSt int
		fmt.Fscan(in, &cntSt)
		for j := 0; j < cntSt; j++ {
			var st int
			fmt.Fscan(in, &st)
			stations[i][st] = true
		}
	}

	intersection := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			for k := range stations[i] {
				if _, ok := stations[j][k]; ok {
					intersection[i] = append(intersection[i], j)
					intersection[j] = append(intersection[j], i)
					break
				}
			}
		}
	}

	var a, b int
	fmt.Fscan(in, &a, &b)

	start := make([]int, 0)
	end := make([]int, 0)
	for i := 0; i < m; i++ {
		for k := range stations[i] {
			if k == a {
				start = append(start, i)
			}
			if k == b {
				end = append(end, i)
			}
		}
	}

	visited := make([]int, m)
	for i := 0; i < m; i++ {
		visited[i] = -1
	}
	bfs(intersection, visited, start)
	res := visited[end[0]]
	for _, v := range end[1:] {
		if visited[v] < res {
			res = visited[v]
		}
	}
	fmt.Fprintln(out, res)
}
