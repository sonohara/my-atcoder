package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// https://atcoder.jp/contests/abs/tasks/abc088_b
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) int {
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	an := []int{}
	sum := 0

	for i := 0; i < n; i++ {
		a := nextInt(sc)
		an = append(an, a)
		sum += a
	}

	sort.Ints(an)

	suma := 0
	for i := n - 1; i >= 0; i -= 2 {
		suma += an[i]
	}

	return suma - (sum - suma)
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
