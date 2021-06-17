package jzoffer

var tmpSlice []int
var reversePairsCnt int

// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/submissions/
func reversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	tmpSlice = make([]int, len(nums))
	reversePairsCnt = 0
	mergeSlice(nums, 0, len(nums)-1)
	return reversePairsCnt
}

func mergeSlice(nums []int, start, end int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	mergeSlice(nums, start, mid)
	mergeSlice(nums, mid+1, end)
	// 因为是从两个开始合并的，所以只计算计算跨区域的逆序数
	tmpIndex := 0
	leftiIdx := start
	rightIdx := mid + 1
	currentRightLessCnt := 0 // 记录到当前右边的值排进tmpslice的个数
	for leftiIdx <= mid && rightIdx <= end {
		if nums[leftiIdx] <= nums[rightIdx] {
			tmpSlice[tmpIndex] = nums[leftiIdx]
			leftiIdx++
			tmpIndex++
			reversePairsCnt += currentRightLessCnt
			continue
		}
		tmpSlice[tmpIndex] = nums[rightIdx]
		tmpIndex++
		rightIdx++
		currentRightLessCnt++
	}
	// 如果左边有剩余，全是逆序
	if leftiIdx <= mid {
		copy(tmpSlice[tmpIndex:], nums[leftiIdx:mid+1])
		reversePairsCnt += (mid - leftiIdx + 1) * currentRightLessCnt
		tmpIndex += mid - leftiIdx + 1
	}
	// 如果右边有剩余，都不是逆序
	if rightIdx <= end {
		copy(tmpSlice[tmpIndex:], nums[rightIdx:end+1])
		tmpIndex += end - rightIdx + 1
	}
	copy(nums[start:], tmpSlice[:tmpIndex])
}
