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
				sc: bufio.NewScanner(strings.NewReader(`4 500 300 900 1300 1700`)),
			},
			want: 1300,
		},
		{
			name: "Example2",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`5 99 100 200 300 400 500`)),
			},
			want: -1,
		},
		{
			name: "Example3",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`4 500 100 600 1100 1600`)),
			},
			want: 600,
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
