package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abs/tasks/arc065_a
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)
	buf := make([]byte, 4096)
	sc.Buffer(buf, int(math.Pow10(5)))

	S := nextLine(sc)

	var s string
	s = S
	for len(s) > 0 {
		found := false
		for _, v := range []string{"dream", "dreamer", "erase", "eraser"} {
			if strings.HasSuffix(s, v) {
				s = strings.TrimSuffix(s, v)
				found = true
				break
			}
		}
		if !found {
			break
		}
	}

	if len(s) == 0 {
		return "YES"
	}

	return "NO"
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
	if err := sc.Err(); err != nil {
		panic(err)
	}
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
