package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	//чтобы не путаться с индексами в задаче счёт с 1, и я буду делать с 1
	data := make([][]int, n+1)
	for i := range data {
		data[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			fmt.Fscan(in, &data[i][j])
		}
	}

	prefixsum := make([][]int, n+1)
	for i := range data {
		prefixsum[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if i-1 >= 1 && j-1 >= 1 {
				prefixsum[i][j] = prefixsum[i-1][j] + prefixsum[i][j-1] -
					prefixsum[i-1][j-1] + data[i][j]
			} else if i-1 >= 1 {
				prefixsum[i][j] = prefixsum[i-1][j] + data[i][j]
			} else if j-1 >= 1 {
				prefixsum[i][j] = prefixsum[i][j-1] + data[i][j]
			} else {
				prefixsum[i][j] = data[i][j]
			}
		}
	}

	fmt.Println(data)
	fmt.Println(prefixsum)
	for i := 0; i < k; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		fmt.Println(x1, y1, x2, y2)
		var res int
		//правый нижний угол
		//минус всё что выше
		//минус всё что левее
		//плюс то что вычли дважды
		fmt.Printf("step one: %d\n", prefixsum[x2][y2]-prefixsum[x2-x1][y2])
		fmt.Printf("step two: %d\n", prefixsum[x2][y2]-prefixsum[x2-x1][y2]-
			prefixsum[x2][y2-y1])
		fmt.Printf("step three: %d\n", prefixsum[x2][y2]-prefixsum[x2-x1][y2]-
			prefixsum[x2][y2-y1]+prefixsum[x1-1][y1-1])
		res += prefixsum[x2][y2] - prefixsum[x1-1][y2] -
			prefixsum[x2][y1-1] + prefixsum[x1-1][y1-1]
		fmt.Println(res)
	}
}
