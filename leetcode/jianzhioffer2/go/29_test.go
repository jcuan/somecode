package jzoffer

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			"2",
			args{
				matrix: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 0},
					{0, 2, 3, 4, 1},
				},
			},
			[]int{1, 2, 3, 4, 5, 0, 1, 4, 3, 2, 0, 6, 7, 8, 9},
		},
		{
			"3",
			args{
				matrix: [][]int{
					{2, 5},
					{8, 4},
					{0, -1},
				},
			},
			[]int{2, 5, 4, -1, 0, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder2(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
