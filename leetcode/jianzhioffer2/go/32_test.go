package jzoffer

import "testing"

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 5}

	res := levelOrder2(root)
	t.Errorf("just show %v", res)
}
