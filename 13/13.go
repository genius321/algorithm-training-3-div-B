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

	stackNum := []int{}
	for {
		b, err := in.ReadByte()
		if err != nil {
			break
		}
		if b >= '0' && b <= '9' {
			stackNum = append(stackNum, int(b-'0'))
		} else if b == '+' || b == '-' || b == '*' {
			c := stackNum[len(stackNum)-1]
			stackNum = stackNum[:len(stackNum)-1]
			a := stackNum[len(stackNum)-1]
			stackNum = stackNum[:len(stackNum)-1]
			switch b {
			case '+':
				stackNum = append(stackNum, a+c)
			case '-':
				stackNum = append(stackNum, a-c)
			case '*':
				stackNum = append(stackNum, a*c)
			}
		}
	}
	fmt.Fprintln(out, stackNum[0])
}
