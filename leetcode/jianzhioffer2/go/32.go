package jzoffer

// 使用非递归地算法，需要后序遍历，因为这样左右才是等同的，否则root出栈之后，不好判断层数
// 因为不知道最大层数，引入了map
func levelOrder(root *TreeNode) [][]int {
	values := map[int][]int{}
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{}
	node := root
	maxLevel := 0 // 从0开始
	addToValues := func(val int) {
		level := len(stack)
		if level > maxLevel {
			maxLevel = level
		}
		values[level] = append(values[level], val)
	}
	var lastVisted *TreeNode
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// 现在开始右子树
		stackLen := len(stack)
		if stackLen == 0 {
			break
		}
		node = stack[stackLen-1] // 这里不出栈
		if node.Right == nil || node.Right == lastVisted {
			stack = stack[:stackLen-1]
			addToValues(node.Val)
			lastVisted = node
			node = nil
		} else { // 说明右边还有东西，还不能访问这个结点
			node = node.Right
		}
	}
	retVals := make([][]int, maxLevel+1, maxLevel+1)
	for i := 0; i <= maxLevel; i++ {
		retVals[i] = values[i]
	}
	return retVals
}

// 层次遍历
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var tempLevelData []int        // 临时保存这一行的数据
	var nextLevelQueue []*TreeNode // 保存下一行的节点信息
	var currentQueue []*TreeNode
	currentQueue = append(currentQueue, root)
	for len(currentQueue) != 0 {
		node := currentQueue[0]
		currentQueue = currentQueue[1:]
		tempLevelData = append(tempLevelData, node.Val)
		if node.Left != nil {
			nextLevelQueue = append(nextLevelQueue, node.Left)
		}
		if node.Right != nil {
			nextLevelQueue = append(nextLevelQueue, node.Right)
		}

		if len(currentQueue) == 0 { // 本层处理完毕，开始处理下一层
			result = append(result, tempLevelData)
			tempLevelData = nil
			currentQueue = nextLevelQueue
			nextLevelQueue = nil
		}
	}
	return result
}
