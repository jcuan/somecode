package leetcode

func DeviDeNum(num int) [][]int {
	var res [][]int
	forDeviDeNum(num, num, &[]int{}, &res)
	return res
}

func forDeviDeNum(n, m int, prefix *[]int, res *[][]int) {
	if n <= 0 || m <= 0 {
		return
	}
	if n == m && n == 1 {
		temp := make([]int, len(*prefix))
		copy(temp, *prefix)
		temp = append(temp, 1)
		*res = append(*res, temp)
		return
	}
	if n == m {
		*prefix = append(*prefix, n)
		temp := make([]int, len(*prefix))
		copy(temp, *prefix)
		*res = append(*res, temp)
		*prefix = (*prefix)[:len(*prefix)-1]

		forDeviDeNum(n, m-1, prefix, res)
		return
	}
	if n < m {
		forDeviDeNum(n, n, prefix, res)
	}
	if n > m {
		// 选择m作为公因式最大值
		*prefix = append(*prefix, m)
		forDeviDeNum(n-m, m, prefix, res)
		*prefix = (*prefix)[:len(*prefix)-1]

		forDeviDeNum(n, m-1, prefix, res)
	}
}

// func forDeviDeNum(num int, prefix *[]int, prefixSum int, prefixMap map[int]bool, res *[][]int) {
// 	if prefixSum == num {
// 		temp := make([]int, len(*prefix))
// 		copy(temp, *prefix)
// 		*res = append(*res, temp)
// 		return
// 	}
// 	for i := num - prefixSum; i > 0; i-- {
// 		if prefixMap[i] == true {
// 			continue
// 		}
// 		*prefix = append(*prefix, i)
// 		prefixSum += i
// 		prefixMap[i] = true
// 		forDeviDeNum(num, prefix, prefixSum, prefixMap, res)
// 		prefixMap[i] = false
// 		prefixSum -= i
// 		*prefix = (*prefix)[:len(*prefix)-1]
// 	}
// }
