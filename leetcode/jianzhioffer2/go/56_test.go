package jzoffer

import (
	"reflect"
	"sort"
	"testing"
)

func Test_singleNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	"basic1",
		// 	args{[]int{2, 1, 2, 3}},
		// 	[]int{1, 3},
		// },
		{
			"basic1",
			args{[]int{4, 1, 4, 6}},
			[]int{1, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumbers(tt.args.nums)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
