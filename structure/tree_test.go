package structure

import (
	"testing"
)

type orderFunc func(node *BiTreeNode) []interface{}

func TestTreeOrder(t *testing.T) {
	var node *BiTreeNode
	node = NewBiTreeNode(1)
	node.Left = NewBiTreeNode(2)
	node.Right = NewBiTreeNode(3)
	node.Left.Left = NewBiTreeNode(4)
	node.Left.Right = NewBiTreeNode(5)

	var res []interface{}
	var ret bool
	var intRes []int

	funcs := []orderFunc{
		PreOrder,
		InOrder,
		PostOrder,
	}
	stackFuncs := []orderFunc{
		PreOrderStack,
		InOrderStack,
		PostOrderStack,
	}
	expects := [][]int{
		{1, 2, 4, 5, 3},
		{4, 2, 5, 1, 3},
		{4, 5, 2, 3, 1},
	}

	for i := range funcs {
		res = funcs[i](node)
		ret, intRes = ValuesCheckInt(res, expects[i])
		if ret != true {
			t.Errorf("not equal, %v %v", intRes, expects[i])
		}
		res = stackFuncs[i](node)
		ret, intRes = ValuesCheckInt(res, expects[i])
		if ret != true {
			t.Errorf("not equal, %v %v", intRes, expects[i])
		}
	}

}
