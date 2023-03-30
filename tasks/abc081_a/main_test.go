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
		want int
	}{
		// {
		// 	name: "1",
		// 	args: args{
		// 		r: strings.NewReader("101"),
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "2",
		// 	args: args{
		// 		r: strings.NewReader("000"),
		// 	},
		// 	want: 0,
		// },
		{
			name: "3",
			args: args{
				r: strings.NewReader("011"),
			},
			want: 2,
		},
		// {
		// 	name: "4",
		// 	args: args{
		// 		r: strings.NewReader("111"),
		// 	},
		// 	want: 3,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.r); got != tt.want {
				t.Errorf("f() = %v, want %v", got, tt.want)
			}
		})
	}
}
