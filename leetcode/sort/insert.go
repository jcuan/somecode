package sort

import "fmt"

// 插入排序

// Shell 希尔排序，从小到大
func Shell(inputs []int) {
	var group int // 组的编号
	var step int  // shell排序的步长，按照5 3 1的方式
	var ilen = len(inputs)

	for step = 5; step >= 1; step -= 2 {
		for group = 0; group < step; group++ {
			// 这里使用直接插入排序，从第二个开始
			// 这里是直接插入排序需要循环的次数，group里的元素除了第一个，都需要循环一次
			// 比如对于M1,...,M(step),...,M(2*step) 需要循环的次数就是2次，可以计算出来
			for memberInGroup := group + step; memberInGroup < ilen; memberInGroup += step {
				// 现在memberInGroup前边的在本group里边的元素已经有序
				temp := inputs[memberInGroup]     // 保存下来，这个位置可能被前边移动来的元素覆盖
				memberNow := memberInGroup - step // 当前组已经排好序的最后一个元素，注意本组内的元素间隔都是step
				for ; memberNow >= 0 && inputs[memberNow] > temp; memberNow -= step {
					// 将这些值往后边移动
					inputs[memberNow+step] = inputs[memberNow]
				}
				// 所有比temp大的值都移动到后面去了，最后被移动的元素的index是memberNow+step，现在把temp放在这个位置即可
				inputs[memberNow+step] = temp
			}
		}
		fmt.Printf("step=%d:%v\n", step, inputs)
	}
}

// BinaryInsert 二分插入 仍需要大量移动
func BinaryInsert(inputs []int) {
	findBinaryFunc := func(index int) int {
		high := index - 1
		low := 0
		var mid int
		for low <= high {
			mid = (high + low) / 2
			if inputs[index] >= inputs[mid] { // 如果相等的话也在原元素后边，保持稳定性
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		return low
	}
	// 需要循环len-1次
	for i := 1; i < len(inputs); i++ {
		current := inputs[i]
		// 已经有序的序列0...i-1 二分查找
		index := findBinaryFunc(i)
		for j := i; j > index; j-- {
			inputs[j] = inputs[j-1]
		}
		inputs[index] = current
	}
}
