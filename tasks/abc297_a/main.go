package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abc297/tasks/abc297_a
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) int {
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	d := nextInt(sc)
	x1 := 0
	x2 := 0

	for i := 0; i < n; i++ {
		x2 = nextInt(sc)
		if x1 == 0 {
			x1 = x2
			continue
		}
		if x2-x1 <= d {
			return x2
		}
		x1 = x2
	}

	return -1
}

// Greatest Common Divisor
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
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
