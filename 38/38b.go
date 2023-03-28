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
			if board[x+dx[i]][y+dy[i]] == -1 {
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
	//добовляю непроходимые клетки, по 2 по краям
	n += 4
	m += 4
	//т.к. нумерация с 0
	s += 1
	t += 1
	fleas := make([][2]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &fleas[i][0], &fleas[i][1])
		fleas[i][0] += 1
		fleas[i][1] += 1
	}
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i < 2 || j < 2 {
				board[i][j] = -2
			} else if i > n-3 || j > m-3 {
				board[i][j] = -2
			} else {
				board[i][j] = -1
			}
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
