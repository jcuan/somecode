package jzoffer

func singleNumbers(nums []int) []int {
	x := 0
	for _, val := range nums {
		x ^= val
	}
	// 找到不为0的那一位
	var index uint
	ux := uint(x)
	for ux > 0 {
		if ux&1 != 0 {
			break
		}
		ux = ux >> 1
		index++
	}
	num1 := 0
	num2 := 0
	for _, val := range nums {
		if uint(val)&(1<<index) == 0 {
			num1 ^= val
		} else {
			num2 ^= val
		}
	}
	return []int{num1, num2}
}
