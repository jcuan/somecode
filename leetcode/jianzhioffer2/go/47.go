package jzoffer

func maxValue0(grid [][]int) int {
	rowcnt := len(grid)
	columncnt := len(grid[0])
	dp := make([][]int, rowcnt)
	for i := 0; i < rowcnt; i++ {
		dp[i] = make([]int, columncnt)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < rowcnt; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < columncnt; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < rowcnt; i++ {
		for j := 1; j < columncnt; j++ {
			if dp[i][j-1] > dp[i-1][j] {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
		}
	}
	return dp[rowcnt-1][columncnt-1]
}

func maxValue(grid [][]int) int {
	rowcnt := len(grid)
	columncnt := len(grid[0])
	dp := make([]int, columncnt)
	dp[0] = grid[0][0]
	for i := 1; i < columncnt; i++ {
		dp[i] = grid[0][i] + dp[i-1]
	}
	for i := 1; i < rowcnt; i++ {
		dp[0] = grid[i][0] + dp[0]
		for j := 1; j < columncnt; j++ {
			if dp[j-1] > dp[j] {
				dp[j] = dp[j-1] + grid[i][j]
			} else {
				dp[j] = dp[j] + grid[i][j]
			}
		}
	}
	return dp[columncnt-1]
}
