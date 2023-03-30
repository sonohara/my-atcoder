package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Print(f(os.Stdin))
}

func f(r io.Reader) string {
	var (
		a int
		b int
		c int
		s string
	)
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	fmt.Fscan(r, &c)
	fmt.Fscan(r, &s)
	return fmt.Sprintf("%d %s", a+b+c, s)
}
