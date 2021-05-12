package nonleetcode

import "testing"

func TestSimpleExpressionCal(t *testing.T) {
	datas := []string{
		"",
		"1",
		"1+1-2+2*4+4/2",
		"11+11-22-10*10+20/20",
		"-1+2",
		"-1+(-2)*(1+1)",
		"+1+1",
		"1-(+1+1)",
	}
	resRights := []int{0, 1, 10, -99, 1, -5, 2, -1}
	for i := range datas {
		res, err := SimpleExpressionCal2(datas[i])
		resRight := resRights[i]
		if resRight == 0 {
			if err != ExpressionError {
				t.Errorf("err case failed\n")
			}
			continue
		}
		if resRight != res {
			t.Errorf("res is wrong, expect:[%v] res:[%v]", resRight, res)
		}
	}

}
