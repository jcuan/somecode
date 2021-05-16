package sort

// Heap 堆排序
func Heap(inputs []int) {
	ilen := len(inputs)
	buildMaxHeap(inputs, ilen)
	for i := ilen - 1; i > 0; i-- {
		inputs[i], inputs[0] = inputs[0], inputs[i] // 交换堆顶元素和堆末元素
		maxHeapify(inputs, 0, i)
	}
}

// https://leetcode-cn.com/problems/sort-an-array/solution/pai-xu-shu-zu-by-leetcode-solution/
func buildMaxHeap(inputs []int, ilen int) {
	// 除以2是因为对于堆的特点，所有可能的根，就是这些
	for i := ilen / 2; i >= 0; i-- {
		maxHeapify(inputs, i, ilen)
	}
}

// 为什么只有遍历一次：因为是在已经建好的堆上调整，除了root，都是满足堆的（包括建立初堆的时候）
func maxHeapify(heap []int, root, heapLen int) {
	nextLeft := root*2 + 1
	for nextLeft < heapLen {
		maxIndex := nextLeft
		nextRight := nextLeft + 1
		if nextRight < heapLen && heap[nextRight] > heap[nextLeft] {
			maxIndex = nextRight
		}
		if heap[maxIndex] > heap[root] {
			heap[maxIndex], heap[root] = heap[root], heap[maxIndex]
			root = maxIndex
			nextLeft = root*2 + 1
		} else {
			break
		}
	}
}
