package jzoffer

// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
func lengthOfLongestSubstring(s string) int {
	indexMap := map[byte]int{} // 记录字符和其对应的位置
	left := 0
	maxcnt := 0
	currentcnt := 0
	for i := 0; i < len(s); i++ {
		if dupIndex, ok := indexMap[s[i]]; ok && dupIndex != -1 { // 说明在map中，出现了重复
			if currentcnt > maxcnt {
				maxcnt = currentcnt
			}
			// 不用删除对应的key，因为字符总数不多，直接标记为-1
			for j := left; j < dupIndex; j++ {
				indexMap[s[j]] = -1
			}
			indexMap[s[i]] = i
			currentcnt = currentcnt - (dupIndex - left)
			left = dupIndex + 1
			continue
		}
		indexMap[s[i]] = i
		currentcnt++
	}
	if currentcnt > maxcnt {
		return currentcnt
	}
	return maxcnt
}
