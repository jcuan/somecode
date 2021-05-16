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
	Val  interface{}
	Next *ListNode
}

// MustIntVal 得到int的值
func (l *ListNode) MustIntVal() int {
	resInt, ok := l.Val.(int)
	if !ok {
		panic("MustIntVal fail!")
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

// 不创建任何新节点
func MergeTwoIntList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var newHead *ListNode
	var pother *ListNode
	if l1.MustIntVal() <= l2.MustIntVal() {
		newHead = l1
		pother = l2
	} else {
		newHead = l2
		pother = l1
	}
	pnew := newHead.Next
	prenew := newHead
	for pnew != nil && pother != nil {
		if pnew.MustIntVal() <= pother.MustIntVal() {
			prenew = pnew
			pnew = pnew.Next
			continue
		}
		prenew.Next = pother
		prenew = prenew.Next
		pother = pother.Next
		prenew.Next = pnew
	}
	if pother != nil {
		prenew.Next = pother
	}
	return newHead
}

// 创建一个新节点
func MergeTwoIntList2(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for l1 != nil && l2 != nil {
		if l1.MustIntVal() <= l2.MustIntVal() {
			cur.Next = l1
			cur = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return newHead.Next
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
		b.WriteString(fmt.Sprintf(flag, p.Val))
		p = p.Next
	}
	log.Println(b.String())
}

// NewIntList 根据int slice创建list
func NewIntList(Vals []int) (*ListNode, error) {
	if len(Vals) == 0 {
		return nil, errors.New("empty Vals")
	}
	var head *ListNode
	var pre *ListNode
	for i := range Vals {
		l := new(ListNode)
		l.Val = Vals[i]
		if head == nil {
			head = l
		} else {
			pre.Next = l
		}
		pre = l
	}
	return head, nil
}
