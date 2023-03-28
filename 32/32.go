package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(vertices, comp [][]int, visited []int, now, numComp int) {
	visited[now] = numComp
	comp[numComp] = append(comp[numComp], now)
	for _, v := range vertices[now] {
		if visited[v] == 0 {
			dfs(vertices, comp, visited, v, numComp)
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
	numComp := 0
	//0 индекс пустует для удобства
	comp := make([][]int, 1)
	for i, l := 1, len(visited); i < l; i++ {
		if visited[i] == 0 {
			numComp++
			comp = append(comp, []int{})
			dfs(vertices, comp, visited, i, numComp)
		}
	}
	fmt.Fprintln(out, numComp)
	for i := 1; i <= numComp; i++ {
		fmt.Fprintln(out, len(comp[i]))
		for _, v := range comp[i] {
			fmt.Fprint(out, v, " ")
		}
		fmt.Fprintln(out)
	}
}
