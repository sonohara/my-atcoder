package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abs/tasks/abc087_b
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) int {
	sc.Split(bufio.ScanWords)

	a := nextInt(sc)
	b := nextInt(sc)
	c := nextInt(sc)
	x := nextInt(sc)
	ans := 0

	for _a := 0; _a <= a; _a++ {
		for _b := 0; _b <= b; _b++ {
			for _c := 0; _c <= c; _c++ {
				if _a*500+_b*100+_c*50 == x {
					ans++
				}
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
