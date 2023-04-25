package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abc070/tasks/abc070_b
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) int {
	sc.Split(bufio.ScanWords)
	// デフォルトのバッファサイズより入力長が大きい場合、正しく読み込めないため拡張しておく
	buf := make([]byte, 4096)
	sc.Buffer(buf, int(math.Pow10(5)))

	a := float64(nextInt(sc))
	b := float64(nextInt(sc))
	c := float64(nextInt(sc))
	d := float64(nextInt(sc))

	return int(math.Max(0, math.Min(b, d)-math.Max(a, c)))
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
