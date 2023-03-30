package main

import (
	"fmt"
	"io"
	"os"
)

// https://atcoder.jp/contests/abs/tasks/abc086_a
func main() {
	fmt.Print(f(os.Stdin))
}

func f(r io.Reader) string {
	var (
		a int
		b int
	)
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)

	if (a*b)%2 == 0 {
		return "Even"
	}

	return "Odd"
}
