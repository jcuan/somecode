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
	var p1, p2, result, resultEnd *ListNode //resultEnd指向返回结果的最后一个有效node
	p1, result = l1, l1                     //将l2的加到l1上来
	p2 = l2
	value := 0 //每一位数加上来的值
	carry := 0 //进位信息
	for p1 != nil && p2 != nil {
		value = p1.Val + p2.Val + carry
		if value > 9 {
			carry = 1
			value %= 10
		} else {
			carry = 0
		}
		p1.Val = value
		resultEnd = p1
		p1 = p1.Next
		p2 = p2.Next
	}
	var plast *ListNode
	if p1 == nil {
		plast = p2
	} else {
		plast = p1
	}
	if plast == nil {
		if carry == 1 {
			addLastNode(resultEnd)
		}
		return result
	}
	if carry == 0 {
		resultEnd.Next = plast
		return result
	} else {
		resultEnd.Next = plast
		value = plast.Val + carry
		carry = 0
		if value <= 9 { //如果没有进位，直接返回
			plast.Val = value
			return result
		}
		for value > 9 {
			plast.Val = 0
			carry = 1
			resultEnd = plast
			plast = plast.Next
			if plast == nil {
				break
			}
			value = plast.Val + carry
			carry = 0
		}
		//从循环出来的时候，不要忘记把最后value的值赋值给plast
		if plast != nil {
			plast.Val = value
			return result
		}
		if carry == 1 { //需要创建新的
			addLastNode(resultEnd)
		}
		return result
	}
}

//现在两个链表已经加完毕，有进位需要申请一个新的node
func addLastNode(plast *ListNode) {
	newNode := &ListNode{1, nil}
	plast.Next = newNode
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
