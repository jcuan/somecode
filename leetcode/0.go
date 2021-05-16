package leetcode

import (
	"sort"
)

// 递归
func Permute46(nums []int) [][]int {
	var res [][]int
	forPermute46(nums, 0, &res)
	return res
}

// 递归 去重复
func Permute47(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	forPermute47(nums, 0, &res)
	return res
}

// 回溯法
func PermuteTraceBack46(nums []int) [][]int {
	var res [][]int
	llen := len(nums)
	if llen == 0 {
		return res
	}
	used := make([]bool, llen) // 默认值是false
	prefix := []int{}
	dfs46(nums, 0, &prefix, used, &res)
	return res
}

// // 回溯法
// func PermuteTraceBack47(nums []int) [][]int {
// 	var res [][]int
// 	llen := len(nums)
// 	if llen == 0 {
// 		return res
// 	}
// 	sort.Ints(nums)
// 	used := make([]bool, llen) // 默认值是false
// 	prefix := []int{}
// 	dfs47(nums, 0, &prefix, used, &res)
// 	return res
// }

// func dfs47(nums []int, first int, prefix *[]int, used []bool, res *[][]int) {
// 	nlen := len(nums)
// 	if first == nlen {
// 		temp := make([]int, nlen)
// 		copy(temp, *prefix)
// 		*res = append(*res, temp)
// 		return
// 	}
// 	for i := 0; i < nlen; i++ {
// 		if used[i] == true {
// 			continue
// 		}
// 		if i > first && nums
// 		used[i] = true
// 		*prefix = append(*prefix, nums[i])
// 		dfs46(nums, first+1, prefix, used, res)
// 		used[i] = false
// 		*prefix = (*prefix)[:len(*prefix)-1]
// 	}
// }

func dfs46(nums []int, first int, prefix *[]int, used []bool, res *[][]int) {
	nlen := len(nums)
	if first == nlen {
		temp := make([]int, nlen)
		copy(temp, *prefix)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < nlen; i++ {
		if used[i] == true {
			continue
		}
		used[i] = true
		*prefix = append(*prefix, nums[i])
		dfs46(nums, first+1, prefix, used, res)
		used[i] = false
		*prefix = (*prefix)[:len(*prefix)-1]
	}
}

func forPermute46(nums []int, first int, res *[][]int) {
	nlen := len(nums)
	if first == nlen-1 {
		tmp := make([]int, nlen)
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i := first; i < nlen; i++ {
		nums[first], nums[i] = nums[i], nums[first]
		forPermute46(nums, first+1, res)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

func forPermute47(nums []int, first int, res *[][]int) {
	nlen := len(nums)
	if first == nlen-1 {
		tmp := make([]int, nlen)
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i := first; i < nlen; i++ {
		if first != i && nums[i] == nums[first] {
			continue
		}
		nums[first], nums[i] = nums[i], nums[first]
		forPermute47(nums, first+1, res)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

// 动态规划，未优化
// https://link.zhihu.com/?target=https%3A//leetcode-cn.com/problems/unique-paths/
func uniquePaths62(m int, n int) int {
	var info = make([][]int, m)
	for i := 0; i < m; i++ {
		info[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		info[i][0] = 1
	}
	for i := 0; i < n; i++ {
		info[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			info[i][j] = info[i-1][j] + info[i][j-1]
		}
	}
	return info[m-1][n-1]
}

// 空间优化版本
func minPathSum62_2(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	var info = make([]int, n)
	info[0] = grid[0][0]
	for i := 1; i < n; i++ {
		info[i] = info[i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		info[0] = info[0] + grid[i][0]
		for j := 1; j < n; j++ {
			var temp int
			if info[j] > info[j-1] {
				temp = info[j-1]
			} else {
				temp = info[j]
			}
			info[j] = temp + grid[i][j]
		}
	}
	return info[n-1]
}

// 动态规划，未优化
func uniquePathsWithObstacles64(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	var info = make([][]int, m)
	for i := 0; i < m; i++ {
		info[i] = make([]int, n)
	}
	// 注意这个的边界条件是个坑，一旦出现一个障碍，所有的都将不可达
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] != 0 {
			info[i][0] = -1
			for j := i + 1; j < m; j++ {
				info[j][0] = -1
			}
			break
		} else {
			info[i][0] = 1
		}
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] != 0 {
			info[0][i] = -1
			for j := i + 1; j < n; j++ {
				info[0][j] = -1
			}
			break
		} else {
			info[0][i] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 0 {
				info[i][j] = -1
			} else {
				found := false
				if info[i-1][j] != -1 {
					info[i][j] += info[i-1][j]
					found = true
				}
				if info[i][j-1] != -1 {
					info[i][j] += info[i][j-1]
					found = true
				}
				if found == false {
					info[i][j] = -1
				}
			}
		}
	}
	if info[m-1][n-1] == -1 {
		return 0
	}
	return info[m-1][n-1]
}

func minPathSum64(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	var info = make([][]int, m)
	for i := 0; i < m; i++ {
		info[i] = make([]int, n)
	}
	info[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		info[i][0] = info[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		info[0][i] = info[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			var temp int
			if info[i-1][j] > info[i][j-1] {
				temp = info[i][j-1]
			} else {
				temp = info[i-1][j]
			}
			info[i][j] = temp + grid[i][j]
		}
	}
	return info[m-1][n-1]
}

func minDistance76(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	var info = make([][]int, len1+1) // 注意需要多一个处理空字符串
	for i := 0; i <= len1; i++ {
		info[i] = make([]int, len2+1)
	}
	// 初始条件
	for i := 0; i <= len1; i++ {
		info[i][0] = i
	}
	for i := 0; i <= len2; i++ {
		info[0][i] = i
	}
	var temp1, temp2, temp3, res int
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				info[i][j] = info[i-1][j-1]
			} else {
				temp1 = info[i-1][j-1] + 1 // 替换
				temp2 = info[i-1][j] + 1   // 删除
				temp3 = info[i][j-1] + 1   // 增加一个
				if temp1 < temp2 {
					res = temp1
				} else {
					res = temp2
				}
				if res > temp3 {
					res = temp3
				}
				info[i][j] = res
			}
		}
	}
	return info[len1][len2]
}

// 优化内存使用版本
// | angle | up
// | left | x=?
// |
// 每一个的计算涉及到三个值，如果使用一维j+1个变量，需要把这个temp保存下来，并且对info[0]特殊处理，等于此时的i值
func minDistanceMem76(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	var info = make([]int, len2+1) // 注意需要多一个处理空字符串
	for i := 0; i <= len2; i++ {
		info[i] = i
	}
	var angle int
	var temp1, temp2, temp3, res int
	for i := 1; i <= len1; i++ {
		angle = info[0]
		info[0] = i
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				info[j], angle = angle, info[j]
			} else {
				temp1 = angle + 1     // 替换
				temp2 = info[j] + 1   // 删除
				temp3 = info[j-1] + 1 // 增加一个
				if temp1 < temp2 {
					res = temp1
				} else {
					res = temp2
				}
				if res > temp3 {
					res = temp3
				}
				angle = info[j]
				info[j] = res
			}
		}
	}
	return info[len2]
}

// 中心拓展算法
func longestPalindromeMiddle5(s string) string {
	slen := len(s)
	if slen <= 1 {
		return s
	}
	middleFunc := func(start, end int) int {
		left, right := start, end
		for left >= 0 && right < slen && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1 // 实际是末尾-开始+1 不过左右各多走了一个位置，所以最终的长度这么算
	}

	var currentLen int
	var temp int
	var start, end int
	for i := 0; i < slen-1; i++ {
		len1 := middleFunc(i, i)   // 单中心
		len2 := middleFunc(i, i+1) // 双中心
		if len1 > len2 {
			temp = len1
		} else {
			temp = len2
		}
		if currentLen < temp {
			currentLen = temp
			start = i - (currentLen-1)/2
			end = i + currentLen/2
		}
	}
	return s[start : end+1]
}
