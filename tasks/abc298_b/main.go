package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abc298/tasks/abc298_b
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)

	a := make([][]int, n)
	b := make([][]int, n)

	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = nextInt(sc)
		}
	}

	for i := 0; i < n; i++ {
		b[i] = make([]int, n)
		for j := 0; j < n; j++ {
			b[i][j] = nextInt(sc)
		}
	}

	x := 0
	match := true
	_a := make([][]int, n)
	copy(_a, a)
	for x <= 3 {
		y := 0
		for y < x {
			for i := 0; i < n; i++ {
				_a[i] = make([]int, n)
				for j := 0; j < n; j++ {
					_a[i][j] = a[n-1-j][i]
				}
			}
			y++
		}
		copy(a, _a)
		match = true
	loop:
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if _a[i][j] == 1 && b[i][j] != 1 {
					match = false
					break loop
				}
			}
		}
		if match {
			break
		}
		x++
	}

	if match {
		return "Yes"
	} else {
		return "No"
	}
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
