package jzoffer

// https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
/*
	1. 找到所在的范围[start, end] 比如[10, 99], [100, 999]
	2. 找到所在范围的那一个数targetNum
	3. 找到这个数的数位
*/
func findNthDigit(n int) int {
	val := n
	digit := 1 // 所在范围每个数占用的位数
	cnt := 9   // 对应的个数
	for val > 0 {
		val -= digit * cnt
		digit++
		cnt *= 10
	}
	// 还原值
	digit--
	cnt /= 10
	val += digit * cnt
	beforeNum := val / digit
	reminder := val % digit
	// 现在对应的区间start
	start := 1
	for i := digit; i > 1; i-- {
		start *= 10
	}
	targetNum := start + beforeNum - 1
	if reminder != 0 {
		targetNum++
	} else {
		reminder = digit // 用于后续k的计算，没有余数的可以看做reminder为digit
	}
	// 将这个targetNum进行分解，需要拿到最高位，也就是除以10*(digit-reminder），然后取得%10
	k := digit - reminder
	for k > 0 {
		targetNum /= 10
		k--
	}
	return targetNum % 10
}

func findNthDigit1(n int) int {
	val := n
	digit := 1               // 所在范围每个数占用的位数
	cnt := 9                 // 对应的数字个数
	weishuCnt := digit * cnt // 对应的位数个数
	start := 1
	for val > weishuCnt {
		val -= weishuCnt
		digit++
		cnt *= 10
		start *= 10
		weishuCnt = digit * cnt
	}
	beforeNum := val / digit
	reminder := val % digit
	// 现在对应的区间start
	targetNum := start + beforeNum - 1
	if reminder != 0 {
		targetNum++
	} else {
		reminder = digit // 用于后续k的计算，没有余数的可以看做reminder为digit
	}
	// 将这个targetNum进行分解，需要拿到最高位，也就是除以10*(digit-reminder），然后取得%10
	k := digit - reminder
	for k > 0 {
		targetNum /= 10
		k--
	}
	return targetNum % 10
}
