package structure

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	var node *BiTreeNode
	node = NewBiTreeNode('A')
	node.Left = NewBiTreeNode('B')
	node.Right = NewBiTreeNode('C')
	node.Left.Left = NewBiTreeNode('D')
	node.Left.Right = NewBiTreeNode('E')

	fmt.Println("PreOrder")
	PreOrder(node)
	fmt.Println("")
	PreOrderStack(node)
	fmt.Println("")
	fmt.Println("MiddleOrder")
	MiddleOrder(node)
	fmt.Println("")
	MiddleOrderStack(node)
	fmt.Println("")
	fmt.Println("PostOrder")
	PostOrder(node)
	fmt.Println("")
	PostOrderStack(node)
	fmt.Println("")
}
