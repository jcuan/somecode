package jzoffer

import (
	"reflect"
	"sort"
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	"basic1",
		// 	args{[]int{3, 4, 5, 1, 2}, 3},
		// 	[]int{1, 2, 3},
		// },
		{
			"basic2",
			args{[]int{3, 2, 1}, 2},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLeastNumbers(tt.args.arr, tt.args.k)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
