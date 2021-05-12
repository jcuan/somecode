package leetcode

import "github.com/jcuan/somecode/structure"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree106(inorder []int, postorder []int) *structure.BiTreeNode {
	length := len(inorder) // 两个顺序一定是一样长
	var node *structure.BiTreeNode
	if length == 0 {
		return nil
	}
	if length == 1 {
		node = &structure.BiTreeNode{Val: inorder[0]}
		return node
	}
	rootVal := postorder[length-1]
	node = &structure.BiTreeNode{Val: rootVal}
	index := findIndexInInorder(inorder, rootVal)
	node.Left = buildTree106(inorder[0:index], postorder[0:index])
	if index+1 < length {
		node.Right = buildTree106(inorder[index+1:], postorder[index:length-1])
	} else {
		node.Right = nil
	}
	return node
}

func findIndexInInorder(inorder []int, find int) int {
	ilen := len(inorder)
	for i := 0; i < ilen; i++ {
		if inorder[i] == find {
			return i
		}
	}
	return -1
}
