package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abc296/tasks/abc296_c
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	x := nextInt(sc)
	if x == 0 {
		return "Yes"
	}

	a := map[int]int{}
	ax := []int{}
	for i := 0; i < n; i++ {
		_a := nextInt(sc)
		ax = append(ax, _a+x)
		a[_a] = _a
	}
	for i := 0; i < n; i++ {
		_, ok := a[ax[i]]
		if ok {
			return "Yes"
		}
	}
	return "No"
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
