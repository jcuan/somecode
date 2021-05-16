package jzoffer

func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		// 圈上
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		// 圈右， 第一个被占用
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		// 圈下，第一个被占用
		if top < bottom {
			for column := right - 1; column >= left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
		}
		// 圈左，第一个和最后一个被占用
		if left < right {
			for row := bottom - 1; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

func spiralOrder(matrix [][]int) []int {
	var res []int
	xcnt := len(matrix)
	if xcnt == 0 {
		return res
	}
	ycnt := len(matrix[0])
	turn := 0
	xvalid := xcnt // 用来指示是否还可以转圈
	yvalid := ycnt
	for xvalid >= 2 && yvalid >= 2 {
		xvalid -= 2 // 用来指示是否还可以转圈
		yvalid -= 2
		// 圈上
		for i := turn; i < ycnt-turn; i++ {
			res = append(res, matrix[turn][i])
		}
		// 圈右，注意第一个被圈上已经访问了
		for i := turn + 1; i < xcnt-turn; i++ {
			res = append(res, matrix[i][ycnt-turn-1])
		}
		// 圈下，注意第一个被圈右
		for i := ycnt - turn - 2; i >= turn; i-- {
			res = append(res, matrix[xcnt-turn-1][i])
		}
		// 圈左，注意第一个被圈下，最后一个被圈上使用
		for i := xcnt - turn - 2; i > turn; i-- {
			res = append(res, matrix[i][turn])
		}
		turn += 1
	}

	if xvalid == 0 || yvalid == 0 {
		return res
	}
	xMatched := false
	if xvalid >= yvalid {
		xMatched = true
		for i := turn; i <= xcnt-turn-1; i++ {
			res = append(res, matrix[i][turn])
		}
	}
	if !xMatched {
		for i := turn; i <= ycnt-turn-1; i++ {
			res = append(res, matrix[turn][i])
		}
	}
	return res
}
