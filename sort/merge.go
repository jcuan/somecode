package sort

import (
	st "github.com/jcuan/somecode/structure"
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

var tmpForMergeSlice []int

// slice版本的归并排序
func MergeSortSlice(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	tmpForMergeSlice = make([]int, len(data))
	MergeSlice(data, 0, len(data)-1)
	return data
}

// 使用一个全局变量代替每次调用MergeSlice append数据，可以从80ms -> 60ms左右
func MergeSlice(data []int, left, right int) {
	if right-left <= 1 { // 特殊处理，加快
		if data[left] > data[right] {
			data[left], data[right] = data[right], data[left]
		}
		return
	}
	// 递归
	middle := (left + right) / 2
	// 不断地进行左右对半划分
	MergeSlice(data, left, middle)
	MergeSlice(data, middle+1, right)
	// 合并
	i := left
	j := middle + 1
	tmpidx := 0
	for i <= middle && j <= right {
		if data[i] <= data[j] {
			tmpForMergeSlice[tmpidx] = data[i]
			tmpidx++
			i++
			continue
		}
		tmpForMergeSlice[tmpidx] = data[j]
		tmpidx++
		j++
	}
	// 有剩下的, append不可能越界，不会有内存重新分配和地址变化
	if i <= middle {
		copy(tmpForMergeSlice[tmpidx:], data[i:middle+1])
		tmpidx += middle - i + 1
	} else if j <= right {
		copy(tmpForMergeSlice[tmpidx:], data[j:right+1])
		tmpidx += right - j + 1
	}
	copy(data[left:], tmpForMergeSlice[:tmpidx])
}
