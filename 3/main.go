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
	subStringCharactors := [256]int{}
	result := 0
	current := 0         //当前子串的上次
	subStart := 0        //当前子串的开始在s中的位置
	duplicatedIndex := 0 //子串中重复的字符的位置+1
	var value uint8
	for index, charactor := range s {
		value = uint8(charactor)
		duplicatedIndex = subStringCharactors[value]
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
