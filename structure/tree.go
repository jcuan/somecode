package structure

import "fmt"

// BiTreeNode 二叉树结点
type BiTreeNode struct {
	Val   interface{}
	Left  *BiTreeNode
	Right *BiTreeNode
}

// PrintFunc 打印函数，可以自定义
var PrintFunc = func(val interface{}) {
	fmt.Printf(" %c ", val)
}

// NewBiTreeNode 新结点
func NewBiTreeNode(val interface{}) *BiTreeNode {
	return &BiTreeNode{Val: val}
}

// MiddleOrder 中序遍历
func MiddleOrder(node *BiTreeNode) {
	if node != nil {
		MiddleOrder(node.Left)
		PrintFunc(node.Val)
		MiddleOrder(node.Right)
	}
}

// MiddleOrderStack 非递归通过栈的方式
func MiddleOrderStack(root *BiTreeNode) {
	node := root
	s1 := NewStack()
	var temp *BiTreeNode
	for node != nil || !s1.IsEmpty() {
		for node != nil {
			s1.Push(node)
			node = node.Left
		}
		val, _ := s1.Pop()
		temp, _ = val.(*BiTreeNode)
		PrintFunc(temp.Val)
		node = temp.Right
	}
}

// PreOrder 先序遍历
func PreOrder(node *BiTreeNode) {
	if node != nil {
		PrintFunc(node.Val)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// PreOrderStack 非递归通过栈的方式
func PreOrderStack(root *BiTreeNode) {
	node := root
	s1 := NewStack()
	var temp *BiTreeNode
	for node != nil || !s1.IsEmpty() {
		for node != nil {
			PrintFunc(node.Val)
			s1.Push(node)
			node = node.Left
		}
		val, _ := s1.Pop()
		temp, _ = val.(*BiTreeNode)
		node = temp.Right
	}
}

// PostOrder 后序遍历
func PostOrder(node *BiTreeNode) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		PrintFunc(node.Val)
	}
}

// PostOrderStack 非递归通过栈的方式
func PostOrderStack(root *BiTreeNode) {
	node := root
	s1 := NewStack()
	var temp *BiTreeNode
	var lastVisted *BiTreeNode
	for node != nil || !s1.IsEmpty() {
		for node != nil {
			s1.Push(node)
			node = node.Left
		}
		val, _ := s1.Top() // 注意这个地方不需要出栈，因为遍历右子树回来的时候还需要这个结点
		temp, _ = val.(*BiTreeNode)
		// 必须右子树为空或者刚刚被访问过才可以访问该根节点，否则访问右子树
		if temp.Right == nil || temp.Right == lastVisted {
			PrintFunc(temp.Val)
			s1.Pop() // 这个结点没有用处啦
			lastVisted = temp
			node = nil // 两种情况下这个node都不再使用，需要使用栈里边的了
		} else {
			node = temp.Right
		}
	}
}
