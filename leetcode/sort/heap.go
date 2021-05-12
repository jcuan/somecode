package sort

// Heap 堆排序
func Heap(inputs []int) {
	ilen := len(inputs)
	buildHeap(inputs, ilen)
	for i := ilen - 1; i > 0; i-- {
		inputs[i], inputs[0] = inputs[0], inputs[i] // 交换堆顶元素和堆末元素
		updateHeap(inputs, 0, i)
	}
}

// buildHeap 建堆 大顶堆
func buildHeap(inputs []int, ilen int) {
	start := (ilen - 2) / 2
	// 标号从0开始取就这关系，这个可以通过下边的关系推出来，不管最后一个有没有右子树，都是要从那一个开始的，所有都用右子树的算start
	// 还有函数关系式的改变，从1开始编号 y=2*k (结点k和它的左子树编号值得关系)
	// 先在从0开始编号，等式成立：y-1=2*(k-1), y=2*(k-1)+1=2*t+1 （新编号情况下t和它的左子树编号关系）
	for i := start; i >= 0; i-- {
		downLeft := 2*i + 1 // 标号从0开始取得话就是这个关系
		downRight := downLeft + 1
		maxIndex := i
		if inputs[downLeft] > inputs[maxIndex] { // downleft不可能超出范围
			maxIndex = downLeft
		}
		if downRight < ilen && inputs[downRight] > inputs[maxIndex] {
			maxIndex = downRight
		}
		// 需要循环交换
		if maxIndex != i {
			inputs[i], inputs[maxIndex] = inputs[maxIndex], inputs[i]
			// 被交换的那个，一直往后看是否需要循环交换
			updateHeap(inputs, maxIndex, ilen)
		}
	}
}

// 更新了节点index，需要向下排查是否需要更新
func updateHeap(inputs []int, index, ilen int) {
	currentChanged := index // 记录当前被移动的元素，需要向下追溯的
	maxIndex := index
	downLeft := 2*currentChanged + 1
	downRight := downLeft + 1
	for downLeft < ilen {
		if inputs[downLeft] > inputs[maxIndex] { // downleft不可能超出范围
			maxIndex = downLeft
		}
		if downRight < ilen && inputs[downRight] > inputs[maxIndex] {
			maxIndex = downRight
		}
		if currentChanged == maxIndex {
			// 发现并不需要交换，提前退出，追溯完毕，不需要向下了
			break
		}
		inputs[maxIndex], inputs[currentChanged] = inputs[currentChanged], inputs[maxIndex]
		currentChanged = maxIndex
		downLeft = 2*currentChanged + 1
		downRight = downLeft + 1
	}
}
