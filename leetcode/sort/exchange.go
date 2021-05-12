package sort

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
		privot := partition(inputs, low, high)
		Quik(inputs, low, privot-1)
		Quik(inputs, privot+1, high)
	}
}

func partition(inputs []int, low, high int) int {
	var (
		pivot = inputs[low] // 就选择第一个吧
		left  = low
		right = high
	)

	// 注意,循环里是从右侧开始的,最后停下来的地方一定是应该放在low这段的
	for left != right {
		for left < right && pivot < inputs[right] {
			right--
		}

		for left < right && pivot >= inputs[left] { // 相等的也放左边
			left++
		}

		if left < right {
			inputs[left], inputs[right] = inputs[right], inputs[left]
		}
	}
	// 需要把low的值赋值给现在的left,因为是以left为新的中间点
	inputs[low], inputs[left] = inputs[left], inputs[low]

	return left
}
