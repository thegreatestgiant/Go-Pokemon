package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "First Test",
			args: args{
				"hello world",
			},
			want: []string{
				"hello",
				"world",
			},
		},
		{
			name: "Cap",
			args: args{"Hello World!"},
			want: []string{
				"hello",
				"world!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
