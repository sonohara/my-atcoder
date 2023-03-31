package main

import (
	"bufio"
	"strings"
	"testing"
)

func Test_f(t *testing.T) {
	type args struct {
		sc *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				sc: bufio.NewScanner(strings.NewReader("3 8 12 40")),
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				sc: bufio.NewScanner(strings.NewReader("5 5 6 8 10")),
			},
			want: 0,
		},
		{
			name: "Example3",
			args: args{
				sc: bufio.NewScanner(strings.NewReader("6 382253568 723152896 37802240 379425024 404894720 471526144")),
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.sc); got != tt.want {
				t.Errorf("f() = %v, want %v", got, tt.want)
			}
		})
	}
}
