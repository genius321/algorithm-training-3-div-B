package main

import (
	"bufio"
	"fmt"
	"os"
)

var dx = [8]int{1, 1, -1, -1, 2, 2, -2, -2}
var dy = [8]int{2, -2, 2, -2, 1, -1, 1, -1}

func bfs(board [][]int, s, t, n, m int) {
	board[s][t] = 0
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{s, t})
	for len(queue) > 0 {
		x := queue[0][0]
		y := queue[0][1]
		for i, l := 0, len(dx); i < l; i++ {
			if x+dx[i] >= 1 && x+dx[i] <= n &&
				y+dy[i] >= 1 && y+dy[i] <= m &&
				board[x+dx[i]][y+dy[i]] == -1 {
				board[x+dx[i]][y+dy[i]] = board[x][y] + 1
				queue = append(queue, [2]int{x + dx[i], y + dy[i]})
			}
		}
		queue = queue[1:]
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, s, t, q int
	fmt.Fscan(in, &n, &m, &s, &t, &q)
	fleas := make([][2]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &fleas[i][0], &fleas[i][1])
	}
	board := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		board[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			board[i][j] = -1
		}
	}
	bfs(board, s, t, n, m)
	var res int
	for _, v := range fleas {
		tmp := board[v[0]][v[1]]
		if tmp == -1 {
			fmt.Fprintln(out, -1)
			return
		}
		res += tmp
	}
	fmt.Fprintln(out, res)
}
