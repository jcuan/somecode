package jzoffer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var valInorderIndexMap map[int]int
var orderPre []int
var orderIn []int

func buildTree3(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	valInorderIndexMap = make(map[int]int, len(preorder))
	for i := range inorder {
		valInorderIndexMap[inorder[i]] = i
	}
	orderPre = preorder
	orderIn = inorder
	tree := forBuildTree3(0, 0, len(inorder))
	return tree
}

// 在中序中，root的index将中序划分为两部分
// 两个顺序的左子树：[left, ..., rootInOrderIndex)
// 两个顺序的右子树：[rootInOrderIndex+1, ..., right)
func forBuildTree3(rootPreorderIndex, left, right int) *TreeNode {
	rootobj := TreeNode{Val: orderPre[rootPreorderIndex]}
	rootInOrderIndex := valInorderIndexMap[rootobj.Val]
	if left < rootInOrderIndex { // 说明还有左子树
		rootobj.Left = forBuildTree3(rootPreorderIndex+1, left, rootInOrderIndex)
	}
	if rootInOrderIndex+1 < right {
		// root在中序的位置需要换算，等于root在中序的位置+左子树结点数+1
		rootobj.Right = forBuildTree3(rootPreorderIndex+(rootInOrderIndex-left)+1, rootInOrderIndex+1, right)
	}
	return &rootobj
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	valInorderIndexMap = make(map[int]int, len(preorder))
	for i := range inorder {
		valInorderIndexMap[inorder[i]] = i
	}
	root := forBuildTree2(preorder, inorder, 0)
	return root
}

// inorderStartIndex用来指定inorder从原始inorder slice开始的位置
func forBuildTree2(preorder []int, inorder []int, inorderStartIndex int) *TreeNode {
	// preorder的第一个就是根节点
	nodeNum := len(preorder)
	if nodeNum == 0 {
		return nil
	}
	root := TreeNode{Val: preorder[0]}
	// 少递归一次
	indexInorder := valInorderIndexMap[root.Val] - inorderStartIndex
	// 左子树是否还有
	if indexInorder != 0 { // 还有左子树
		root.Left = forBuildTree2(preorder[1:indexInorder+1], inorder[:indexInorder], inorderStartIndex) // 一直到这里都是左子树的范围
	}
	if indexInorder != nodeNum-1 { // 还有右子树
		root.Right = forBuildTree2(preorder[indexInorder+1:nodeNum], inorder[indexInorder+1:nodeNum], valInorderIndexMap[root.Val]+1) // 注意这一不要使用indexInorder+1了
	}
	return &root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// preorder的第一个就是根节点
	nodeNum := len(preorder)
	if nodeNum == 0 {
		return nil
	}
	root := TreeNode{Val: preorder[0]}
	// 少递归一次
	indexInorder := findIndex(root.Val, inorder)
	// 左子树是否还有
	if indexInorder != 0 { // 说明没有左子树了, root在中序也是第一个
		root.Left = buildTree(preorder[1:indexInorder+1], inorder[:indexInorder]) // 一直到这里都是左子树的范围
	}
	if indexInorder != nodeNum-1 { // 没有右子树了
		root.Right = buildTree(preorder[indexInorder+1:nodeNum], inorder[indexInorder+1:nodeNum])
	}
	return &root
}

func findIndex(val int, nodes []int) int {
	for i := range nodes {
		if nodes[i] == val {
			return i
		}
	}
	return -1
}
