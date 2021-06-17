package jzoffer

import "math/rand"

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/submissions/
func getLeastNumbers(arr []int, k int) []int {
	quickSelect(arr, k, 0, len(arr)-1)
	return arr[0:k]
}

// start不能移动，只能为0
func quickSelect(arr []int, k, start, end int) {
	for start <= end {
		mid := partition(arr, k, start, end)
		if mid == k-1 {
			return
		}
		if mid < k-1 {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
}

func partition(arr []int, k, start, end int) int {
	index := rand.Intn(end+1-start) + start
	arr[index], arr[end] = arr[end], arr[index]

	retIndex := start
	for j := start; j < end; j++ {
		if arr[j] <= arr[end] {
			arr[retIndex], arr[j] = arr[j], arr[retIndex]
			retIndex++
		}
	}

	arr[retIndex], arr[end] = arr[end], arr[retIndex]
	return retIndex
}
