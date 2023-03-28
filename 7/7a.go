package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, c string
	fmt.Fscan(in, &a, &b, &c)

	a1 := strings.Split(a, ":")
	aH, _ := strconv.Atoi(a1[0])
	aM, _ := strconv.Atoi(a1[1])
	aS, _ := strconv.Atoi(a1[2])
	aS = aH*60*60 + aM*60 + aS

	b1 := strings.Split(b, ":")
	bH, _ := strconv.Atoi(b1[0])
	bM, _ := strconv.Atoi(b1[1])
	bS, _ := strconv.Atoi(b1[2])
	bS = bH*60*60 + bM*60 + bS

	c1 := strings.Split(c, ":")
	cH, _ := strconv.Atoi(c1[0])
	cM, _ := strconv.Atoi(c1[1])
	cS, _ := strconv.Atoi(c1[2])
	cS = cH*60*60 + cM*60 + cS

	if cS < aS {
		cS += 24 * 60 * 60
	}
	diff := cS - aS
	if diff%2 == 1 {
		diff += 1
	}
	diff /= 2
	resS := bS + diff
	if resS >= 24*60*60 {
		resS -= 24 * 60 * 60
	}

	res := []byte{}
	resH := resS / (60 * 60)
	resM := resS/60 - resH*60
	resS = resS % 60

	if resH == 0 {
		res = append(res, []byte{'0', '0'}...)
	} else if resH > 0 && resH <= 9 {
		res = append(res, []byte{'0', byte(resH) + '0'}...)
	} else {
		res = append(res, strconv.Itoa(resH)...)
	}
	res = append(res, ':')
	if resM == 0 {
		res = append(res, []byte{'0', '0'}...)
	} else if resM > 0 && resM <= 9 {
		res = append(res, []byte{'0', byte(resM) + '0'}...)
	} else {
		res = append(res, strconv.Itoa(resM)...)
	}
	res = append(res, ':')
	if resS == 0 {
		res = append(res, []byte{'0', '0'}...)
	} else if resS > 0 && resS <= 9 {
		res = append(res, []byte{'0', byte(resS) + '0'}...)
	} else {
		res = append(res, strconv.Itoa(resS)...)
	}
	fmt.Fprintln(out, string(res))
}
