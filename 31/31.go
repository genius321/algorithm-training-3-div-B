package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(vertices [][]int, visited []bool, now int) {
	visited[now] = true
	for _, v := range vertices[now] {
		if !visited[v] {
			dfs(vertices, visited, v)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	//0 не трогаю, он просто есть для удобства
	vertices := make([][]int, n+1)
	//создаю список смежности
	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Fscan(in, &v1, &v2)
		vertices[v1] = append(vertices[v1], v2)
		vertices[v2] = append(vertices[v2], v1)
	}

	visited := make([]bool, n+1)
	dfs(vertices, visited, 1)
	var countVertices int
	for _, v := range visited {
		if v == true {
			countVertices++
		}
	}
	fmt.Fprintln(out, countVertices)

	for i, v := range visited {
		if v {
			fmt.Fprint(out, i, " ")
		}
	}
	fmt.Fprintln(out)
}
