package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type symbol struct {
	ch  byte
	cnt int
	pos int
}

type symbols []symbol

func (a symbols) Len() int           { return len(a) }
func (a symbols) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a symbols) Less(i, j int) bool { return a[i].ch < a[j].ch }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	m := make(map[byte]int)
	for {
		ch, err := in.ReadByte()
		if err != nil {
			break
		}
		if ch >= 33 && ch <= 126 {
			m[ch]++
		}
	}

	var rows int
	l := len(m)
	symbols := make(symbols, l)
	i := 0
	for k, v := range m {
		if v > rows {
			rows = v
		}
		symbols[i].ch = k
		symbols[i].cnt = v
		i++
	}
	sort.Sort(symbols)
	for ; rows > 0; rows-- {
		tmp := make([]byte, l)
		for i, v := range symbols {
			if v.cnt == rows {
				symbols[i].cnt--
				tmp[i] = byte('#')
			} else {
				tmp[i] = byte(' ')
			}
		}
		fmt.Fprintln(out, string(tmp))
	}
	tmp := make([]byte, l)
	for i, v := range symbols {
		tmp[i] = v.ch
	}
	fmt.Fprintln(out, string(tmp))
}
