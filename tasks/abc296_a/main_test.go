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
				sc: bufio.NewScanner(strings.NewReader("6 MFMFMF")),
			},
			want: "Yes",
		},
		{
			name: "Example2",
			args: args{
				sc: bufio.NewScanner(strings.NewReader("9 FMFMMFMFM")),
			},
			want: "No",
		},
		{
			name: "Example3",
			args: args{
				sc: bufio.NewScanner(strings.NewReader("1 F")),
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
