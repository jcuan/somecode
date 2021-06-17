package jzoffer

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"basic1",
			args{[]int{2, 7, 11, 15}, 9},
			[]int{2, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findContinuousSequence(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {
		// 	"basic1",
		// 	args{9},
		// 	[][]int{{2, 3, 4}, {4, 5}},
		// },
		{
			"basic2",
			args{10},
			[][]int{{2, 3, 4}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContinuousSequence(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findContinuousSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
