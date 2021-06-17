package jzoffer

import "testing"

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	"basic_1",
		// 	args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
		// 	"0123456789",
		// },
		{
			"basic",
			args{[]int{121, 12}},
			"12112",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
