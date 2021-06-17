package jzoffer

func translateNum0(num int) int {
	numWeis := []int{}
	val := num
	for val > 0 {
		numWeis = append(numWeis, val%10)
		val /= 10
	}
	// dp[0]代表空串，对应的翻译方法数为0
	prepre := 1 // 代表空字符串
	pre := 1    // 第一个位置
	for i := 2; i <= len(numWeis); i++ {
		index := len(numWeis) - i // 对应的在numWeis中的位置，比如1234的数组，numWeis中保存的是4321，要到倒着来
		current := 0
		if numWeis[index+1] != 0 && numWeis[index+1]*10+numWeis[index] <= 25 { // 说明和前一位组合起来小于25，
			current += prepre
		}
		current += pre
		prepre, pre = pre, current
	}
	return pre
}

// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
func translateNum(num int) int {
	val := num
	// dp[0]代表空串，对应的翻译方法数为0
	prepre := 1 // 代表空字符串
	pre := 1    // 第一个位置
	for val >= 10 {
		gewei := val % 10
		shiwei := (val / 10) % 10
		current := 0
		if shiwei != 0 && shiwei*10+gewei <= 25 { // 说明和前一位组合起来小于25，
			current += prepre
		}
		current += pre
		prepre, pre = pre, current
		val = val / 10
	}
	return pre
}
