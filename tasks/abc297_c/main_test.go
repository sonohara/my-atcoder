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
				sc: bufio.NewScanner(strings.NewReader(`2 3 TTT T.T`)),
			},
			want: "PCT\nT.T",
		},
		{
			name: "Example2",
			args: args{
				sc: bufio.NewScanner(strings.NewReader(`3 5
TTT..
.TTT.
TTTTT
`)),
			},
			want: `PCT..
.PCT.
PCPCT
`,
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
