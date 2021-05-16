package leetcode

import "testing"

func TestMinPathSum62(t *testing.T) {
	grid := [][]int{
		{1, 2},
		{5, 6},
		{1, 1},
	}
	res := minPathSum62_2(grid)
	if res != 8 {
		t.Error("error")
	}
}

// func TestBuildTree106(t *testing.T) {
// 	inorder := []int{9, 3, 15, 20, 7}
// 	postorder := []int{9, 15, 7, 20, 3}
// 	root := buildTree106(inorder, postorder)
// 	structure.PrintFunc = func(v interface{}) {
// 		fmt.Printf(" %d ", v)
// 	}
// 	fmt.Println("")
// 	structure.MiddleOrderStack(root)
// 	fmt.Println("")
// 	structure.PostOrderStack(root)
// }
