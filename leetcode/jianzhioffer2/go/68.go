package jzoffer

// 因为节点可能是自己的祖先，需要把节点自己也加入进去
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pParents := []*TreeNode{}
	qParents := []*TreeNode{}
	node := root
	stack := []*TreeNode{}
	var lastVisted *TreeNode
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if len(stack) == 0 {
			break
		}
		stackLen := len(stack)
		node = stack[stackLen-1]
		if node.Right == nil || lastVisted == node.Right {
			if node.Val == p.Val {
				pParents = make([]*TreeNode, stackLen, stackLen)
				copy(pParents, stack)
				if len(qParents) != 0 {
					break
				}
			} else if node.Val == q.Val {
				qParents = make([]*TreeNode, stackLen, stackLen)
				copy(qParents, stack)
				if len(pParents) != 0 {
					break
				}
			}
			lastVisted = node
			stack = stack[:stackLen-1]
			node = nil
		} else {
			node = node.Right
		}
	}
	// 就先循环吧, 把q的搞成一个map
	qMap := map[int]bool{}
	for i := range qParents {
		qMap[qParents[i].Val] = true
	}
	pLen := len(pParents)
	for i := pLen - 1; i >= 0; i-- {
		if qMap[pParents[i].Val] == true {
			return pParents[i]
		}
	}
	return root
}
