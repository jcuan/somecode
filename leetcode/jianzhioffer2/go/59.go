package jzoffer

type Queue struct {
	Vals []int // 使用固定长度，不用处理队列的增长，循环队列，本使用场景永远不会满
	Cap  int
	len  int
	head int // q.Head() == q.Vals[q.head%q.Cap]
	tail int // q.Tail() == q.Vals[(q.tail-1+q.Cap)%q.Cap]，也就是说tail指向的是队列尾的下一个
}

func (q *Queue) Push(val int) {
	q.Vals[q.tail] = val
	q.tail = (q.tail + 1) % q.Cap
	q.len++
}

func (q *Queue) Pop() int {
	val := q.Vals[q.head]
	q.head = (q.head + 1) % q.Cap
	q.len--
	return val
}

func (q *Queue) PopTail() int {
	lastIndex := (q.tail - 1 + q.Cap) % q.Cap
	val := q.Vals[lastIndex]
	q.tail = lastIndex
	q.len--
	return val
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Tail() int {
	return q.Vals[(q.tail-1+q.Cap)%q.Cap]
}

func (q *Queue) Head() int {
	return q.Vals[q.head%q.Cap]
}

func maxSlidingWindowUseQueueObj(nums []int, k int) []int {
	numcnt := len(nums)
	if k > numcnt || k == 0 {
		return []int{}
	}
	res := make([]int, 0, numcnt-k+1) // 一次申请需要的空间
	queueVals := make([]int, k)       // 循环队列，不用判断是否满或者空，都不会出现
	q := &Queue{Vals: queueVals, Cap: k}
	// 初始化第一个窗口的数据
	for i := 0; i < k; i++ {
		for q.Len() > 0 && q.Tail() < nums[i] { // 清除所有小于该值的
			q.PopTail()
		}
		q.Push(nums[i])
	}
	// 已经形成了窗口
	for i := k; i < numcnt; i++ {
		res = append(res, q.Head())
		// 先出元素
		if nums[i-k] == q.Head() {
			q.Pop()
		}
		// 处理新元素
		for q.Len() > 0 && q.Tail() < nums[i] {
			q.PopTail()
		}
		q.Push(nums[i])
	}
	res = append(res, q.Head())
	return res
}

// https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/submissions/
func maxSlidingWindow(nums []int, k int) []int {
	numcnt := len(nums)
	if k > numcnt || k == 0 {
		return []int{}
	}
	res := make([]int, 0, numcnt-k+1) // 一次申请需要的空间
	queue := make([]int, k)           // 循环队列，不用判断是否满，因为qhead==qtail的时候，表明队列空
	qhead := 0
	qtail := 0 // 指向的是尾巴index的下一个
	for i := 0; i < numcnt; i++ {
		// 先出元素
		if i >= k && nums[i-k] == queue[qhead] {
			qhead = (qhead + 1) % k
		}
		// 处理新元素
		for qhead != qtail {
			tailIndex := (qtail - 1 + k) % k
			if queue[tailIndex] >= nums[i] {
				break
			}
			qtail = tailIndex
		}
		queue[qtail] = nums[i]
		qtail = (qtail + 1) % k
		if i >= k-1 {
			res = append(res, queue[qhead])
		}
	}
	return res
}
