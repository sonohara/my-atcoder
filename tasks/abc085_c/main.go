package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abs/tasks/abc085_c
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)

	N := nextInt(sc)
	Y := nextInt(sc)
	ans := "-1 -1 -1"

loop:
	for i := Y / 10000; i >= 0; i-- {
		y := Y - (i * 10000)
		for j := y / 5000; j >= 0; j-- {
			y := y - (5000 * j)
			if y%1000 == 0 && i+j+(y/1000) == N {
				ans = fmt.Sprintf("%d %d %d", i, j, y/1000)
				break loop
			}
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
