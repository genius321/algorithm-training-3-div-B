package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func bfs(cave [][][]int, z, x, y, n int) {
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{z, x, y})
	for len(queue) > 0 {
		z := queue[0][0]
		x := queue[0][1]
		y := queue[0][2]
		if z+1 < n && cave[z+1][x][y] == -1 {
			cave[z+1][x][y] = cave[z][x][y] + 1
			queue = append(queue, [3]int{z + 1, x, y})
		}
		if z-1 >= 0 && cave[z-1][x][y] == -1 {
			cave[z-1][x][y] = cave[z][x][y] + 1
			queue = append(queue, [3]int{z - 1, x, y})
		}
		if x+1 < n && cave[z][x+1][y] == -1 {
			cave[z][x+1][y] = cave[z][x][y] + 1
			queue = append(queue, [3]int{z, x + 1, y})
		}
		if x-1 >= 0 && cave[z][x-1][y] == -1 {
			cave[z][x-1][y] = cave[z][x][y] + 1
			queue = append(queue, [3]int{z, x - 1, y})
		}
		if y+1 < n && cave[z][x][y+1] == -1 {
			cave[z][x][y+1] = cave[z][x][y] + 1
			queue = append(queue, [3]int{z, x, y + 1})
		}
		if y-1 >= 0 && cave[z][x][y-1] == -1 {
			cave[z][x][y-1] = cave[z][x][y] + 1
			queue = append(queue, [3]int{z, x, y - 1})
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

	cave := make([][][]int, n)
	for i := 0; i < n; i++ {
		cave[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			cave[i][j] = make([]int, n)
		}
	}
	var z, x, y int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				var tmp rune
				fmt.Fscanf(in, "%c", &tmp)
				for !unicode.IsPrint(tmp) {
					fmt.Fscanf(in, "%c", &tmp)
				}
				switch tmp {
				case '#':
					cave[i][j][k] = -2
				case '.':
					cave[i][j][k] = -1
				case 'S':
					cave[i][j][k] = 0
					z, x, y = i, j, k
				}
			}
		}
	}
	if z == 0 {
		fmt.Fprintln(out, 0)
		return
	}
	res := 27000
	bfs(cave, z, x, y, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp := cave[0][i][j]
			if tmp > -1 && tmp < res {
				res = tmp
			}
		}
	}
	fmt.Fprintln(out, res)
}
