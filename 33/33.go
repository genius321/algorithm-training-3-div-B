package main

import (
	"bufio"
	"fmt"
	"os"
)

var ok bool = true

func dfs(vertices [][]int, visited []int, now, color int) {
	visited[now] = color
	for _, v := range vertices[now] {
		if visited[v] == color {
			ok = false
		} else if visited[v] == 0 {
			dfs(vertices, visited, v, 3-color)
		}
		if !ok {
			return
		}
	}
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
		vertices[v2] = append(vertices[v2], v1)
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
	if ok {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
