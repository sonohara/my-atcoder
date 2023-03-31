package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abs/tasks/abc081_b
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) int {
	sc.Split(bufio.ScanWords)

	var n = nextInt(sc)
	var alist []int
	for i := 0; i < n; i++ {
		a := nextInt(sc)
		if a%2 == 1 {
			return 0
		}
		alist = append(alist, a)
	}

	var r int
loop:
	for {
		r++
		for i := 0; i < n; i++ {
			alist[i] = alist[i] / 2
			if alist[i]%2 == 1 {
				break loop
			}
		}
	}

	return r
}

func nextLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextInt(sc *bufio.Scanner) int {
	return atoi(nextLine(sc))
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

var sc = bufio.NewScanner(os.Stdin)
