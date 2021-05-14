package jzoffer

func forIsSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return A.Val == B.Val && forIsSubStructure(A.Left, B.Left) && forIsSubStructure(A.Right, B.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return forIsSubStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// -----
func forIsSubStructure0(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if !(A.Val == B.Val && forIsSubStructure(A.Left, B.Left) && forIsSubStructure(A.Right, B.Right)) {
		if subRoot == B {
			return forIsSubStructure0(A.Left, B) || forIsSubStructure0(A.Right, B)
		}
	}
	return false
}

var subRoot *TreeNode

func isSubStructure0(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	subRoot = B
	return forIsSubStructure(A, B)
}
