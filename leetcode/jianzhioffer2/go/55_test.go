package jzoffer

import "testing"

func TestIsBalanced(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	var res bool
	res = isBalanced(root)
	if res != true {
		t.Error("failed")
	}
	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Right = &TreeNode{Val: 4}
	res = isBalanced(root)
	if res != false {
		t.Error("failed")
	}
}
