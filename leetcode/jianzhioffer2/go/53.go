package jzoffer

const (
	FindTypeLeft = iota
	FindTypeRight
)

func binaryFind(nums []int, target, findType int) int {
	if findType == FindTypeRight {
		if nums[len(nums)-1] == target {
			return len(nums) - 1
		}
	} else {
		if nums[0] == target {
			return 0
		}
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if findType == FindTypeLeft { // findleft：如果有解，一定存在{小于target，target}
				if nums[mid-1] < target {
					return mid
				} else {
					right = mid - 1
				}
			} else { // findright：如果有解，一定存在{target，大于target}
				if nums[mid+1] > target {
					return mid
				} else {
					left = mid + 1
				}
			}
			continue
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	// 分别找左右边界
	// 1. 排除明显不存在的
	if nums[0] > target || nums[len(nums)-1] < target {
		return 0
	}
	// 2. 二分搜索找左边界，找到的条件是，左边的数比target小，右边的数等于target
	left := binaryFind(nums, target, FindTypeLeft)
	if left == -1 {
		return 0
	}
	right := binaryFind(nums, target, FindTypeRight)
	if left == -1 {
		return 0
	}
	return right - left + 1
}
