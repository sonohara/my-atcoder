package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abc297/tasks/abc297_b
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)

	b1 := 0
	b2 := 0
	k := 0
	r1 := 0
	r2 := 0
	s := nextLine(sc)
	for i, v := range strings.Split(s, "") {
		if v == "B" {
			if b1 == 0 {
				b1 = i + 1
			} else {
				b2 = i + 1
			}
		}
		if v == "K" {
			k = i + 1
		}
		if v == "R" {
			if r1 == 0 {
				r1 = i + 1
			} else {
				r2 = i + 1
			}
		}
	}

	if (b1+b2)%2 == 0 {
		return "No"
	}

	if r1 > k || k > r2 {
		return "No"
	}

	return "Yes"
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
