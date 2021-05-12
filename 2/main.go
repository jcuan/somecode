package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	nums1 := []int{9}
	nums2 := []int{1, 9, 9, 9, 9, 9, 8, 9, 9, 9}
	l1, l2 := makeList(nums1, nums2)
	l3 := addTwoNumbers(l1, l2)
	printList(l3)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p1, p2, resHead, pres, newNode *ListNode
	resHead = &ListNode{}
	pres = resHead
	p1, p2 = l1, l2 //分别指向两个链表
	value := 0      //每一位数加上来的值
	carry := 0      //进位信息
	for p1 != nil || p2 != nil {
		value = carry
		if p1 != nil {
			value += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			value += p2.Val
			p2 = p2.Next
		}
		if value > 9 {
			carry = 1
			value %= 10
		} else {
			carry = 0
		}
		newNode = &ListNode{}
		newNode.Next = nil
		newNode.Val = value
		pres.Next = newNode
		pres = pres.Next
	}
	if carry != 0 {
		newNode = &ListNode{}
		newNode.Next = nil
		newNode.Val = carry
		pres.Next = newNode
		pres = pres.Next
	}
	resHead = resHead.Next
	return resHead
}

func makeList(l1, l2 []int) (*ListNode, *ListNode) {
	var temp *ListNode
	head1 := &ListNode{l1[0], nil}
	p1 := head1
	head2 := &ListNode{l2[0], nil}
	p2 := head2
	len1 := len(l1)
	len2 := len(l2)
	for i := 1; i < len1; i++ {
		temp = &ListNode{l1[i], nil}
		p1.Next = temp
		p1 = temp
	}
	for i := 1; i < len2; i++ {
		temp = &ListNode{l2[i], nil}
		p2.Next = temp
		p2 = temp
	}
	return head1, head2
}

func printList(l *ListNode) {
	if l == nil {
		fmt.Println("[]")
		return
	}
	var buffer bytes.Buffer
	buffer.WriteString("[" + strconv.Itoa(l.Val))
	l = l.Next
	for l != nil {
		buffer.WriteString("," + strconv.Itoa(l.Val))
		l = l.Next
	}
	buffer.WriteString("]")
	str := buffer.String()
	fmt.Println(str)
}
