package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abc075/tasks/abc075_b
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)
	// デフォルトのバッファサイズより入力長が大きい場合、正しく読み込めないため拡張しておく
	buf := make([]byte, 4096)
	sc.Buffer(buf, int(math.Pow10(5)))

	H := nextInt(sc)
	W := nextInt(sc)
	S := map[int]map[int]string{}
	c := map[int]map[int]string{}
	T := map[int]map[int]string{}

	for i := 1; i <= H; i++ {
		S[i] = map[int]string{}
		c[i] = map[int]string{}
		s := nextLine(sc)
		for j, v := range strings.Split(s, "") {
			if v == "." {
				c[i][j+1] = ""
			}
			S[i][j+1] = v
		}
	}

	for i, v := range c {
		T[i] = map[int]string{}
		for j := range v {
			t := 0
			if v, ok := S[i-1][j-1]; ok {
				if v == "#" {
					t++
				}
			}
			if v, ok := S[i-1][j]; ok {
				if v == "#" {
					t++
				}
			}
			if v, ok := S[i-1][j+1]; ok {
				if v == "#" {
					t++
				}
			}
			if v, ok := S[i][j-1]; ok {
				if v == "#" {
					t++
				}
			}
			if v, ok := S[i][j+1]; ok {
				if v == "#" {
					t++
				}
			}
			if v, ok := S[i+1][j-1]; ok {
				if v == "#" {
					t++
				}
			}
			if v, ok := S[i+1][j]; ok {
				if v == "#" {
					t++
				}
			}
			if v, ok := S[i+1][j+1]; ok {
				if v == "#" {
					t++
				}
			}
			T[i][j] = itoa(t)
		}
	}

	for i, v := range T {
		for j := range v {
			S[i][j] = T[i][j]
		}
	}

	// fmt.Println(S)
	// fmt.Println(T)

	s := ""
	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			s += S[i][j]
		}
		s += "\n"
	}

	return s
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
