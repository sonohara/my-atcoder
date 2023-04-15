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
				sc: bufio.NewScanner(strings.NewReader(`3
				0 1 1
				1 0 0
				0 1 0
				1 1 0
				0 0 1
				1 1 1
				`)),
			},
			want: "Yes",
		},
		{
			name: "Example2",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`2
				0 0
				0 0
				1 1
				1 1
				`)),
			},
			want: "Yes",
		},
		{
			name: "Example3",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`5
				0 0 1 1 0
				1 0 0 1 0
				0 0 1 0 1
				0 1 0 1 0
				0 1 0 0 1
				1 1 0 0 1
				0 1 1 1 0
				0 0 1 1 1
				1 0 1 0 1
				1 1 0 1 0
				`)),
			},
			want: "No",
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
