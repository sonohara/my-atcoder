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
				sc: bufio.NewScanner(strings.NewReader(`RNBQKBNR`)),
			},
			want: "Yes",
		},
		{
			name: "Example2",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`KRRBBNNQ`)),
			},
			want: "No",
		},
		{
			name: "Example3",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`BRKRBQNN`)),
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
