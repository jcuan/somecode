package sort

import (
	st "github.com/jcuan/leetcode/somecode"
)

// 归并排序

// MergeList 单链表版本归并排序 需要传入链表的大小
func MergeList(l *st.ListNode, llen int) *st.ListNode {
	if llen <= 1 {
		return l
	}
	// 将链表拆分成两个
	middle := llen / 2
	var p *st.ListNode = l
	var left, right *st.ListNode
	for i := 0; i < middle-1; i++ {
		p = p.Next
	}
	left = l
	right = p.Next
	p.Next = nil
	left = MergeList(left, middle)
	right = MergeList(right, llen-middle)
	// 将两个从小到大的有序链表合并成一个
	res := st.MergeTwoIntList(left, right)
	return res
}

// MergeSlice slice版本的归并排序
func MergeSlice(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	// 递归
	middle := len(data) / 2
	// 不断地进行左右对半划分
	left := MergeSlice(data[:middle])
	right := MergeSlice(data[middle:])
	// 合并
	return forMergeSlice(left, right)
}

// MergeSortIntSlice使用
// 这个函数可以说是代价很高了，先这样吧，可读性好
func forMergeSlice(left, right []int) (result []int) {
	l, r := 0, 0 // 对应左边和右边处理的元素的index
	llen := len(left)
	rlen := len(right)
	for l < llen && r < rlen {
		// 从小到大排序
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	// 剩下的处理
	if l < llen {
		result = append(result, left[l:]...)
	}
	if r < rlen {
		result = append(result, right[r:]...)
	}
	return result
}
