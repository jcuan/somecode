package jzoffer

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"basic1",
			args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
			[]int{3, 3, 5, 5, 6, 7},
		},
		{
			"basic_equal",
			args{[]int{3, 1, 3, 2, 1}, 3},
			[]int{3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
