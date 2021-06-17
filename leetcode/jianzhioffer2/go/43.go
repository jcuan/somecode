package jzoffer

// https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/solution/mian-shi-ti-43-1n-zheng-shu-zhong-1-chu-xian-de-2/
func countDigitOne(n int) int {
	val := n
	digit_weight := 1
	res := 0
	for val != 0 {
		reminder := val % 10
		switch {
		case reminder == 0:
			res += (val / 10) * digit_weight
		case reminder == 1:
			res += (val/10)*digit_weight + ((n - val*digit_weight) + 1)
		default:
			res += (val/10 + 1) * digit_weight
		}
		digit_weight *= 10
		val /= 10
	}
	return res
}
