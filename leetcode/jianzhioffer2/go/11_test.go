package jzoffer

import "testing"

func TestExist(t *testing.T) {
	data := [][]byte{
		{'A', 'B', 'C', 'D'},
		{'B', 'G', 'C', 'J'},
		{'F', 'B', 'H', 'I'},
	}
	testcase := map[string]bool{
		"A":     true,
		"H":     true,
		"AB":    true,
		"ABFB":  true,
		"ABFBF": false,
	}
	for str, res := range testcase {
		testdata := make([][]byte, len(data[0])-1)
		for i := range data {
			testdata[i] = make([]byte, len(data[0])-1)
			copy(testdata[i], data[i])
		}
		if exist(testdata, str) != res {
			t.Errorf("not equal, str:%s", str)
		}
	}
}
