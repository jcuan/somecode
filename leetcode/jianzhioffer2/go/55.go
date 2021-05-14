package jzoffer

// tall返回-1代表已经检测到失衡
func getBalanceInfo(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftTall int
	var rightTall int
	if root.Left != nil {
		leftTall = getBalanceInfo(root.Left)
	}
	if root.Right != nil {
		rightTall = getBalanceInfo(root.Right)
	}
	if leftTall == -1 || rightTall == -1 {
		return -1
	}
	var bigger int
	var smaller int
	if leftTall > rightTall {
		bigger = leftTall
		smaller = rightTall
	} else {
		bigger = rightTall
		smaller = leftTall
	}
	if bigger-smaller > 1 {
		return -1
	}
	return bigger + 1
}

// 采用中序的遍历方式
// 同一采用左边减右边的方式
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := getBalanceInfo(root)
	return res != -1
}

func getDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftTall int
	var rightTall int
	if root.Left != nil {
		leftTall = getDeep(root.Left)
	}
	if root.Right != nil {
		rightTall = getDeep(root.Right)
	}
	if leftTall > rightTall {
		return leftTall + 1
	}
	return rightTall + 1
}

func maxDepth(root *TreeNode) int {
	return getDeep(root)
}
