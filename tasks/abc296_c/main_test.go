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
		want string
	}{
		{
			name: "Example1",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`6 5 3 1 4 1 5 9`)),
			},
			want: "Yes",
		},
		{
			name: "Example2",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`6 -4 -2 -7 -1 -8 -2 -8`)),
			},
			want: "No",
		},
		{
			name: "Example3",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`2 0 141421356 17320508`)),
			},
			want: "Yes",
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
