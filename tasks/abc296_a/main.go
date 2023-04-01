package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abc296/tasks/abc296_a
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)

	nextLine(sc)
	s := nextLine(sc)

	if len(s) == 1 {
		return "Yes"
	}

	if strings.Contains(s, "MM") || strings.Contains(s, "FF") {
		return "No"
	}

	return "Yes"
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
