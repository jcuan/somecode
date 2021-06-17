package jzoffer

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"basic",
			args{
				"123",
				"123",
			},
			true,
		},
		{
			"basic_dot",
			args{
				"123",
				"12.",
			},
			true,
		},
		{
			"basic_*_1",
			args{
				"123",
				"12*",
			},
			false,
		},
		{
			"basic_*_2",
			args{
				"123",
				"123*",
			},
			true,
		},
		{
			"basic_*_dot",
			args{
				"123",
				"12.*",
			},
			true,
		},
		{
			"basic_*_dot_ignore",
			args{
				"123",
				"123.*",
			},
			true,
		},
		{
			"complex_*_dot",
			args{
				"123333",
				"123.*",
			},
			true,
		},
		{
			"wrong_*_dot",
			args{
				"1233",
				"12333.*",
			},
			false,
		},
		{
			"wrong_1",
			args{
				"aa",
				"a",
			},
			false,
		},
		{
			"wrong_2",
			args{
				"ab",
				".*",
			},
			true,
		},
		{
			"wrong_3",
			args{
				"aaa",
				"ab*a*c*a",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
