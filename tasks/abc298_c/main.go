package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abc298/tasks/abc298_c
func main() {
	fmt.Print(f(sc))
}

func f(sc *bufio.Scanner) string {
	sc.Split(bufio.ScanWords)

	nextInt(sc)
	q := nextInt(sc)
	hako2 := map[int][]int{}
	hako3 := map[int][]int{}
	// ans := []string{}
	// m := make(map[int]bool)

	for x := 0; x < q; x++ {
		y := nextInt(sc)
		if y == 1 {
			i := nextInt(sc)
			j := nextInt(sc)
			// if _, ok := hako2[j]; !ok {
			// 	hako2[j] = []int{}
			// }
			hako2[j] = append(hako2[j], i)
			// if _, ok := hako3[i]; !ok {
			// 	hako3[i] = []int{}
			// }
			dup := false
			for _, v := range hako3[i] {
				if v == j {
					dup = true
					break
				}
			}
			if !dup {
				hako3[i] = append(hako3[i], j)
			}
			continue
		}
		if y == 2 {
			i := nextInt(sc)
			s := hako2[i]
			sort.Ints(s)
			ss := []string{}
			for i := 0; i < len(s); i++ {
				ss = append(ss, itoa(s[i]))
			}
			fmt.Println(strings.Join(ss, " "))
			continue
		}
		if y == 3 {
			i := nextInt(sc)
			s := hako3[i]
			sort.Ints(s)
			ss := []string{}
			for i := 0; i < len(s); i++ {
				ss = append(ss, itoa(s[i]))
			}
			fmt.Println(strings.Join(ss, " "))
			continue
		}
	}
	return "hoge"
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
