package jzoffer

import (
	"testing"
)

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	"basic_lt_10",
		// 	args{9},
		// 	9,
		// },
		{
			"basic_eq_10",
			args{100},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit1(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNthDigit1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit1(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit1() = %v, want %v", got, tt.want)
			}
		})
	}
}
