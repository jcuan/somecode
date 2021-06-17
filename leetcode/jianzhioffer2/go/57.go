package jzoffer

// 如果能找到，返回对应的index，第二个返回值为true
// 如果不能找到，返回比对应的值小的第一个index或者0，第二个返回值为false
func binaryFind57(nums []int, start, end, target int) (int, bool) {
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid, true
		}
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if start == 0 {
		return start, false
	}
	return start - 1, false // start-1一定比target小
}

func twoSumBinary(nums []int, target int) []int {
	// 先缩小右边界
	endLimit, _ := binaryFind57(nums, 0, len(nums)-1, target)
	// 从左到右选择
	for i := 0; i < endLimit; i++ { // 不去等，因为必须要两个
		index, res := binaryFind57(nums, i+1, endLimit, target-nums[i])
		if res {
			return []int{nums[i], nums[index]}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	left := 0
	right, _ := binaryFind57(nums, 0, len(nums)-1, target) // 使用二分法快速缩小右边界
	for left < right {
		val := nums[left] + nums[right]
		if val == target {
			return []int{nums[left], nums[right]}
		}
		if val > target {
			right = right - 1
		} else {
			left = left + 1
		}
	}
	return []int{}
}

func findContinuousSequence(target int) [][]int {
	i := 1
	j := 2 // 取等的区间
	windowSum := 3
	res := [][]int{}
	maxlimit := (target + 1) / 2
	for i < j {
		switch {
		case windowSum < target: // windows过小，右边界移动
			j++
			if j > maxlimit {
				return res
			}
			windowSum += j
		case windowSum > target: // windows过大，左边界移动
			windowSum -= i
			i++
		case windowSum == target: // 相等后保存结果，左边界移动
			item := make([]int, j-i+1)
			for k := i; k <= j; k++ {
				item[k-i] = k
			}
			res = append(res, item)
			windowSum -= i
			i++
		}

	}
	return res
}
