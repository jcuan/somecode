package leetcode

import (
	"fmt"
	"testing"

	"github.com/jcuan/leetcode/structure"
)

func TestBuildTree106(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree106(inorder, postorder)
	structure.PrintFunc = func(v interface{}) {
		fmt.Printf(" %d ", v)
	}
	fmt.Println("")
	structure.MiddleOrderStack(root)
	fmt.Println("")
	structure.PostOrderStack(root)
}
