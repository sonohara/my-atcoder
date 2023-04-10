package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abc297/tasks/abc297_c
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)

	h := nextInt(sc)
	nextInt(sc)
	s := []string{}

	for i := 0; i < h; i++ {
		l := nextLine(sc)
		s = append(s, strings.ReplaceAll(l, "TT", "PC"))
	}

	return strings.Join(s, "\n")
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
