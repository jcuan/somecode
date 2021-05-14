package jzoffer

import "testing"

func TestIsSubStructure(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 4}

	datas := [][]*TreeNode{
		// {root, nil},
		// {nil, nil},
		{root, &TreeNode{Val: 1}},
		{root, root.Left.Left},
		{root, root.Left},
		{root, root2},
	}
	answers := []bool{true, true, true, false}
	for i := range datas {
		res := isSubStructure(datas[i][0], datas[i][1])
		if res != answers[i] {
			t.Errorf("not equal: %v %v %v", i, res, answers[i])
		}
	}
}
