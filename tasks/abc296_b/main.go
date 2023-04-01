package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abc296/tasks/abc296_b
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)

	r := map[int]string{
		7: "h",
		6: "g",
		5: "f",
		4: "e",
		3: "d",
		2: "c",
		1: "b",
		0: "a",
	}
	j := 0
	for {
		s := nextLine(sc)
		i := strings.Index(s, "*")
		if i >= 0 {
			return r[i] + itoa(8-j)
		}
		j++
	}
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
