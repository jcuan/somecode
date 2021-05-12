package structure

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

// ListNode 链表结点
// 注意这个链表不含有头结点
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// MustIntValue 得到int的值
func (l *ListNode) MustIntValue() int {
	resInt, ok := l.Value.(int)
	if !ok {
		panic("MustIntValue fail!")
	}
	return resInt
}

// Skip 跳过指定的位置
func (l *ListNode) Skip(step int) *ListNode {
	var p *ListNode
	p = l
	for ; step > 0; step-- {
		p = p.Next
	}
	return p
}

// MergeTwoIntList 将连个list合并
// 直接在l1上操作，不创建新的list
// order必须是从小到大排列
func MergeTwoIntList(l1 *ListNode, l2 *ListNode) *ListNode {
	pl1 := l1
	pl2 := l2

	var (
		head, // 新链表头
		tail, // 新链表尾
		p *ListNode // 操作新链表的指针
	)

	for pl1 != nil && pl2 != nil {
		if pl1.MustIntValue() <= pl2.MustIntValue() {
			p = pl1
			pl1 = pl1.Next
		} else {
			p = pl2
			pl2 = pl2.Next
		}
		if head == nil { // 中间不处理Next乱指的问题，最后再置空
			head = p
			tail = p
		} else {
			tail.Next = p
			tail = p
		}
	}
	// 好了，现在至少有一个达到末尾了
	if pl1 != nil {
		p = pl1
	} else if pl2 != nil {
		p = pl2
	} else {
		p = nil
	}
	// 把这个尾巴加到末尾
	for p != nil {
		tail.Next = p
		tail = p
		p = p.Next
	}
	tail.Next = nil
	return head
}

// Print 打印链表
func (l *ListNode) Print() {
	p := l
	var flag string
	var b bytes.Buffer
	for p != nil {
		if p == l {
			flag = "%v"
		} else {
			flag = "->%v"
		}
		b.WriteString(fmt.Sprintf(flag, p.Value))
		p = p.Next
	}
	log.Println(b.String())
}

// NewIntList 根据int slice创建list
func NewIntList(values []int) (*ListNode, error) {
	if len(values) == 0 {
		return nil, errors.New("empty values")
	}
	var head *ListNode
	var pre *ListNode
	for i := range values {
		l := new(ListNode)
		l.Value = values[i]
		if head == nil {
			head = l
		} else {
			pre.Next = l
		}
		pre = l
	}
	return head, nil
}
