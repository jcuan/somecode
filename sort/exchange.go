package sort

import "math/rand"

// Bubble 冒泡排序
func Bubble(inputs []int) {
	ilen := len(inputs)
	var changed = true // 用来指明是否有交换，没有交换可以提前结束排序

	for i := 1; i < ilen; i++ { // 外层循环，需要循环len-1次
		if changed == false {
			break
		}
		changed = false
		for j := 0; j < ilen-i; j++ { // 内层循环，在未排序的前段排序
			if inputs[j] > inputs[j+1] {
				inputs[j], inputs[j+1] = inputs[j+1], inputs[j]
				changed = true
			}
		}
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
