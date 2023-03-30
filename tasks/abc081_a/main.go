package main

import (
	"fmt"
	"io"
	"os"
)

// https://atcoder.jp/contests/abs/tasks/abc081_a
func main() {
	fmt.Print(f(os.Stdin))
}

func f(r io.Reader) int {
	var a int
	fmt.Fscanf(r, "%d", &a)
	fmt.Println(a)
	if a == 0 {
		return 0
	}

	if a == 111 {
		return 3
	}

	if a == 110 || a == 101 || a == 11 {
		return 2
	}

	return 1
}
