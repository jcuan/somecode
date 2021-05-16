package jzoffer

// board被修改
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	// 最后需要把这些还原
	return false
}

// 上 下 左 右
func dfs(matrix [][]byte, x, y, start int, word string) bool {
	// 越界或者不匹配
	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) || matrix[x][y] != word[start] {
		return false
	}

	matchNodeBackup := word[start]
	matrix[x][y] = 0 // 标记为0，防止重复走
	if start == len(word)-1 {
		return true
	}
	// 直接判断的方式
	start += 1
	if dfs(matrix, x-1, y, start, word) {
		return true
	}
	if dfs(matrix, x+1, y, start, word) {
		return true
	}
	if dfs(matrix, x, y-1, start, word) {
		return true
	}
	if dfs(matrix, x, y+1, start, word) {
		return true
	}
	// 都不匹配，回退
	start -= 1
	matrix[x][y] = matchNodeBackup
	return false
}
