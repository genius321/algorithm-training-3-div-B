package main

import (
	"bufio"
	"fmt"
	"os"
)

func bfs(board [][]int, s, t, n, m int) {
	board[s][t] = 0
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{s, t})
	for len(queue) > 0 {
		x := queue[0][0]
		y := queue[0][1]
		if x-1 >= 1 && y-2 >= 1 && board[x-1][y-2] == -1 {
			board[x-1][y-2] = board[x][y] + 1
			queue = append(queue, [2]int{x - 1, y - 2})
		}
		if x+1 <= n && y-2 >= 1 && board[x+1][y-2] == -1 {
			board[x+1][y-2] = board[x][y] + 1
			queue = append(queue, [2]int{x + 1, y - 2})
		}
		if x-1 >= 1 && y+2 <= m && board[x-1][y+2] == -1 {
			board[x-1][y+2] = board[x][y] + 1
			queue = append(queue, [2]int{x - 1, y + 2})
		}
		if x+1 <= n && y+2 <= m && board[x+1][y+2] == -1 {
			board[x+1][y+2] = board[x][y] + 1
			queue = append(queue, [2]int{x + 1, y + 2})
		}
		if x-2 >= 1 && y+1 <= m && board[x-2][y+1] == -1 {
			board[x-2][y+1] = board[x][y] + 1
			queue = append(queue, [2]int{x - 2, y + 1})
		}
		if x-2 >= 1 && y-1 >= 1 && board[x-2][y-1] == -1 {
			board[x-2][y-1] = board[x][y] + 1
			queue = append(queue, [2]int{x - 2, y - 1})
		}
		if x+2 <= n && y-1 >= 1 && board[x+2][y-1] == -1 {
			board[x+2][y-1] = board[x][y] + 1
			queue = append(queue, [2]int{x + 2, y - 1})
		}
		if x+2 <= n && y+1 <= m && board[x+2][y+1] == -1 {
			board[x+2][y+1] = board[x][y] + 1
			queue = append(queue, [2]int{x + 2, y + 1})
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
