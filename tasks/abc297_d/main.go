package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abc297/tasks/abc297_d
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) int {
	sc.Split(bufio.ScanWords)

	a := nextInt(sc)
	b := nextInt(sc)

	i := 0
	for a != b {
		if a > b {
			x := 1
			if a-b > b {
				// x = int(math.Floor(float64((a - b) / b)))
				x = (a - b) / b
			}
			a = a - (b * x)
			i = i + x
		} else {
			x := 1
			if b-a > a {
				// x = int(math.Floor(float64((b - a) / a)))
				x = (b - a) / a
			}
			b = b - (a * x)
			i = i + x
		}
	}

	return i
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
