package sort

import "math/rand"

// Bubble 冒泡排序
func Bubble(inputs []int) {
	ilen := len(inputs)
	var lastChangeIndex = ilen - 1 // 判断交换的最后位置，这之后的都是有序的，不用再判断

	for i := 1; i < ilen; i++ { // 外层循环，需要循环len-1次
		tmpLastUpdate := -1
		for j := 0; j < ilen-i && j <= lastChangeIndex; j++ { // 内层循环，在未排序的前段排序
			if inputs[j] > inputs[j+1] {
				inputs[j], inputs[j+1] = inputs[j+1], inputs[j]
				tmpLastUpdate = j
			}
		}
		if tmpLastUpdate == -1 { // 没有需要再交换的了
			break
		}
		lastChangeIndex = tmpLastUpdate
	}
}

// Quik 快速排序
func Quik(inputs []int, low, high int) {
	if low < high {
		privot := randomPartition(inputs, low, high)
		Quik(inputs, low, privot-1)
		Quik(inputs, privot+1, high)
	}
}

func randomPartition(inputs []int, left, right int) int {
	pivot := rand.Intn(right-left) + left
	inputs[pivot], inputs[right] = inputs[right], inputs[pivot]

	retidx := left
	for j := left; j < right; j++ {
		if inputs[j] <= inputs[right] {
			inputs[j], inputs[retidx] = inputs[retidx], inputs[j]
			retidx++
		}
	}
	// retdix现在的值一定大于pivotval也就是inputs[right], 否则会继续++的
	inputs[retidx], inputs[right] = inputs[right], inputs[retidx]

	return retidx
}
