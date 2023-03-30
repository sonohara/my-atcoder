package main

import (
	"io"
	"strings"
	"testing"
)

func Test_f(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				r: strings.NewReader("1 1"),
			},
			want: "Odd",
		},
		{
			name: "2",
			args: args{
				r: strings.NewReader("1 2"),
			},
			want: "Even",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.r); got != tt.want {
				t.Errorf("f() = %v, want %v", got, tt.want)
			}
		})
	}
}
