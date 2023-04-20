package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// https://atcoder.jp/contests/abs/tasks/arc089_a
func main() {
	fmt.Print(f(sc))
}

// see: https://qiita.com/drken/items/fd4e5e3630d0f5859067#%E7%AC%AC-10-%E5%95%8F--abc-086-c---traveling-300-%E7%82%B9
func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)
	// デフォルトのバッファサイズより入力長が大きい場合、正しく読み込めないため拡張しておく
	buf := make([]byte, 4096)
	sc.Buffer(buf, int(math.Pow10(5)))

	n := nextInt(sc)
	t := 0
	x := 0
	y := 0

	for i := 0; i < n; i++ {
		ti := nextInt(sc)
		xi := nextInt(sc)
		yi := nextInt(sc)

		if int(math.Abs(float64(xi)-float64(x)))+int(math.Abs(float64(yi)-float64(y))) > ti-t {
			return "No"
		}

		if ti%2 != (xi+yi)%2 {
			return "No"
		}

		t, x, y = ti, xi, yi
	}

	return "Yes"
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
