package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abc296/tasks/abc296_d
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) int {
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	m := nextInt(sc)

	if m == n {
		return m
	}

	if n*n < m {
		return -1
	}

	r := 0
	for i := n; i > 1; i-- {
		x := m % i
		if x == 0 && m/i <= n {
			return m
		}
		j := int(math.Ceil(float64(m / i)))
		if j > n {
			break
		}
		if r == 0 || x < r {
			r = x
		}
	}

	return m + r
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