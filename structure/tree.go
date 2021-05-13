package structure

// BiTreeNode 二叉树结点
type BiTreeNode struct {
	Val   interface{}
	Left  *BiTreeNode
	Right *BiTreeNode
}

type OrderType int

const (
	ORDER_PRE = iota
	ORDER_IN
	ORDER_POST
)

// NewBiTreeNode 新结点
func NewBiTreeNode(val interface{}) *BiTreeNode {
	return &BiTreeNode{Val: val}
}

func orderInner(orderType OrderType, values *[]interface{}, node *BiTreeNode) {
	if node == nil {
		return
	}
	if orderType == ORDER_PRE {
		*values = append(*values, node.Val)
	}
	orderInner(orderType, values, node.Left)
	if orderType == ORDER_IN {
		*values = append(*values, node.Val)
	}
	orderInner(orderType, values, node.Right)
	if orderType == ORDER_POST {
		*values = append(*values, node.Val)
	}
}

// InOrder 中序遍历，返回value的列表
func InOrder(node *BiTreeNode) []interface{} {
	values := []interface{}{}
	orderInner(ORDER_IN, &values, node)
	return values
}

// PreOrder 先序遍历
func PreOrder(node *BiTreeNode) []interface{} {
	values := []interface{}{}
	orderInner(ORDER_PRE, &values, node)
	return values
}

// PostOrder 后序遍历
func PostOrder(node *BiTreeNode) []interface{} {
	values := []interface{}{}
	orderInner(ORDER_POST, &values, node)
	return values
}

// ValuesCheckInt 检查对应的数据，返回检查结果和解析后的int列表的值
func ValuesCheckInt(values []interface{}, expects []int) (bool, []int) {
	intValues := ValuesToInt(values)
	if len(intValues) != len(expects) {
		return false, intValues
	}
	for i := range intValues {
		if intValues[i] != expects[i] {
			return false, intValues
		}

	}
	return true, intValues
}

func ValuesToInt(values []interface{}) []int {
	intValues := make([]int, 0, len(values))
	for i := range values {
		inValue := values[i].(int)
		intValues = append(intValues, inValue)
	}
	return intValues
}

// orderStackInner 非递归通过栈的方式: 可以实现先序和中序
func orderStackInner(orderType OrderType, root *BiTreeNode) []interface{} {
	node := root
	s1 := NewStack()
	var temp *BiTreeNode
	values := []interface{}{}
	for node != nil || !s1.IsEmpty() {
		for node != nil {
			if orderType == ORDER_PRE {
				values = append(values, node.Val)
			}
			s1.Push(node)
			node = node.Left
		}
		val, err := s1.Pop()
		if err == EmptyError {
			break
		}
		temp, _ = val.(*BiTreeNode)
		if orderType == ORDER_IN {
			values = append(values, temp.Val)
		}
		node = temp.Right
	}
	return values
}

// InOrderStack 非递归通过栈的方式
func InOrderStack(root *BiTreeNode) []interface{} {
	return orderStackInner(ORDER_IN, root)
}

// PreOrderStack 非递归通过栈的方式
func PreOrderStack(root *BiTreeNode) []interface{} {
	return orderStackInner(ORDER_PRE, root)
}

// PostOrderStack 非递归通过栈的方式
func PostOrderStack(root *BiTreeNode) []interface{} {
	node := root
	s1 := NewStack()
	var temp *BiTreeNode
	var lastVisted *BiTreeNode
	values := []interface{}{}
	for node != nil || !s1.IsEmpty() {
		for node != nil {
			s1.Push(node)
			node = node.Left
		}
		val, err := s1.Top() // 注意这个地方不需要出栈，因为遍历右子树回来的时候还需要这个结点
		if err == EmptyError {
			break
		}
		temp, _ = val.(*BiTreeNode)
		// 必须右子树为空或者刚刚被访问过才可以访问该根节点，否则访问右子树
		if temp.Right == nil || temp.Right == lastVisted {
			s1.Pop() // 这个结点没有用处啦
			values = append(values, temp.Val)
			lastVisted = temp
			node = nil // 两种情况下这个node都不再使用，需要使用栈里边的了
		} else {
			node = temp.Right
		}
	}
	return values
}
