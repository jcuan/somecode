package jzoffer

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left.Left.Left = &TreeNode{Val: 6}

	datas := [][]*TreeNode{
		{root, root.Left, root.Left.Right, root.Left},
		{root, root.Left.Left.Left, root.Left.Right, root.Left},
	}
	for i := range datas {
		data := datas[i]
		res := lowestCommonAncestor(data[0], data[1], data[2])
		if res != data[3] {
			t.Errorf("not equal, %v %v", res.Val, data[3].Val)
		}
	}
}
