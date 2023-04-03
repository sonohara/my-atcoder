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
		// {
		// 	name: "Example1",
		// 	args: args{
		// 		sc: bufio.NewScanner(strings.NewReader(`5 7`)),
		// 	},
		// 	want: 8,
		// },
		// {
		// 	name: "Example2",
		// 	args: args{
		// 		sc: bufio.NewScanner(strings.NewReader(`2 5`)),
		// 	},
		// 	want: -1,
		// },
		{
			name: "Example3",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`100000 10000000000`)),
			},
			want: 10000000000,
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
