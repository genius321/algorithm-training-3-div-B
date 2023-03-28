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

	var n int
	fmt.Fscan(in, &n)

	res := make([]int, n)      //конечный город
	stack := make([][2]int, 0) //цена проживания + индекс начального города
	for i := 0; i < n; i++ {   //текущий город
		var tmp int //цена проживания
		fmt.Fscan(in, &tmp)
		if len(stack) == 0 {
			stack = append(stack, [2]int{tmp, i})
			continue
		}
		for l := len(stack) - 1; l >= 0 && stack[l][0] > tmp; l = len(stack) - 1 {
			res[stack[l][1]] = i
			stack = stack[:l]
		}
		stack = append(stack, [2]int{tmp, i})
	}
	for _, v := range stack {
		res[v[1]] = -1
	}
	for _, v := range res {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
}
