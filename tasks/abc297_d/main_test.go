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
				sc: bufio.NewScanner(strings.NewReader(`3 8`)),
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`1234567890 1234567890`)),
			},
			want: 0,
		},
		{
			name: "Example3",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`1597 987`)),
			},
			want: 15,
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
