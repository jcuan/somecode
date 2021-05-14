package jzoffer

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	datas := []string{
		// "",
		"you",
		"you are good!",
		"  you   are good! ",
	}
	trueAns := []string{
		// "",
		"you",
		"good! are you",
		"good! are you",
	}
	for i := range datas {
		res := ReverseWords2(datas[i])
		if res != trueAns[i] {
			t.Errorf("get:[%v] true:[%v]", res, trueAns[i])
		}
	}
}
