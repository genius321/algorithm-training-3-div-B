package main

import (
	"bufio"
	"fmt"
	"os"
)

var ok bool = true
var res = make([]int, 0)

func dfs(vertices [][]int, visited []int, now, color int) {
	visited[now] = color
	for _, v := range vertices[now] {
		if visited[v] == 0 {
			dfs(vertices, visited, v, color)
		} else if visited[v] == 1 {
			ok = false
		}
		if !ok {
			break
		}
	}
	visited[now] = 2
	res = append(res, now)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	vertices := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Fscan(in, &v1, &v2)
		vertices[v1] = append(vertices[v1], v2)
	}
	visited := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if !ok {
			break
		}
		if visited[i] == 0 {
			dfs(vertices, visited, i, 1)
		}
	}
	if !ok {
		fmt.Fprintln(out, -1)
	} else {
		for i := len(res) - 1; i >= 0; i-- {
			fmt.Fprint(out, res[i], " ")
		}
		fmt.Fprintln(out)
	}
}
