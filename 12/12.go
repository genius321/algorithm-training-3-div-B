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

	stack := []byte{}
	ans := "yes"
	for {
		b, err := in.ReadByte()
		if err != nil {
			break
		}
		if b == '[' || b == '{' || b == '(' {
			stack = append(stack, b)
		} else if b == ')' || b == '}' || b == ']' {
			if len(stack) == 0 {
				ans = "no"
				break
			}
			switch {
			case (b == ')' && stack[len(stack)-1] == '(') ||
				(b == ']' && stack[len(stack)-1] == '[') ||
				(b == '}' && stack[len(stack)-1] == '{'):
				stack = stack[:len(stack)-1]
			default:
				ans = "no"
				break
			}
		}
	}
	if len(stack) > 0 {
		ans = "no"
	}
	fmt.Println(ans)
}
