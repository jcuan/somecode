package jzoffer

import "testing"

func Test_reversePairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"basic1",
			args{[]int{7, 4, 5}},
			2,
		},
		{
			"basic2",
			args{[]int{4, 5}},
			0,
		},
		{
			"basic3",
			args{[]int{7, 4, 5, 3, 3}},
			8,
		},
		{
			"wrong1",
			args{[]int{4, 5, 6, 7}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePairs(tt.args.nums); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
