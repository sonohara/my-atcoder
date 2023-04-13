package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abs/tasks/abc083_b
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) int {
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	a := nextInt(sc)
	b := nextInt(sc)
	ans := 0

	for i := 1; i <= n; i++ {
		j := i
		x := 0
		for j > 0 {
			x += (j % 10)
			j = j / 10
		}
		if a <= x && x <= b {
			ans += i
		}
	}

	return ans
}

// Greatest Common Divisor
// ユークリッドの互除法を利用して、2数の最大公約数を求める
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
