package main

import (
	"fmt"
)

func main() {
	test()
}

func test() {
	list := []string{"au"}
	for _, ele := range list {
		fmt.Println(lengthOfLongestSubstring(ele))
	}
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	locations := [256]uint{}
	resLen := 0
	var subStart uint = 1    //开始位置
	var duplicatedIndex uint //子串中重复的字符的位置+1
	var value uint8
	for i, charactor := range s {
		value = uint8(charactor)
		duplicatedIndex = locations[value]
		if duplicatedIndex-1 >= subStart {
			//比较和以前的result
			if current > result {
				result = current
			}

			current = current - (duplicatedIndex - subStart) + 1
			subStringCharactors[value] = index + 1
			subStart = duplicatedIndex

		} else {
			//存储的位置只+1的，因为为了区别map不存在的情况
			subStringCharactors[value] = index + 1
			current++
		}
	}
	if current > result {
		result = current
	}
	return result
}
