package leetcode

import (
	"fmt"
	"testing"
)

func TestPermute46(t *testing.T) {
	inputs := []int{1, 2, 3}
	res := Permute46(inputs)
	fmt.Println(res)
}

func TestPermute47(t *testing.T) {
	inputs := []int{2, 2, 1, 1}
	res := Permute47(inputs)
	fmt.Println(res)
}
func TestPermuteTraceBack46(t *testing.T) {
	inputs := []int{1, 2, 3}
	res := PermuteTraceBack46(inputs)
	fmt.Println(res)
}

func TestMinDistance76(t *testing.T) {
	data := [][]string{
		{"horse", "ros"},
		{"info", "i"},
		{"info", "xdxa"},
		{"info", "infoxxyy"},
		{"info", "nxfpp"},
	}
	res := []int{3, 3, 4, 4, 4}
	for i, _ := range data {
		temp := minDistance76(data[i][0], data[i][1])
		if temp != res[i] {
			t.Errorf("[%s] [%s] not pass, expect:%d, actual:%d", data[i][0], data[i][1], res[i], temp)
		}
		temp = minDistanceMem76(data[i][0], data[i][1])
		if temp != res[i] {
			t.Errorf("[%s] [%s] not pass, expect:%d, actual:%d", data[i][0], data[i][1], res[i], temp)
		}
	}
}

func TestLongestPalindromeMiddle5(t *testing.T) {
	data := [][]string{
		{"a", "a"},
		{"info", "i"},
		{"xbx", "xbx"},
		{"abba", "abba"},
		{"abcddc", "cddc"},
	}
	for _, info := range data {
		res := longestPalindromeMiddle5(info[0])
		if res != info[1] {
			t.Errorf("str[%s] expect:[%s], actual:[%s]\n", info[0], info[1], res)
		}
	}
}

func TestUniquePathsWithObstacles64(t *testing.T) {
	data := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	if 2 != uniquePathsWithObstacles64(data) {
		t.Error("failed!")
	}
	data2 := [][]int{
		{1, 0},
	}
	if 0 != uniquePathsWithObstacles64(data2) {
		t.Error("failed!")
	}
}
