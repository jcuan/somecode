package main

import (
	"fmt"
)

func main() {
	test()
}

func test() {
	INT32_MAX := 1<<31 - 1
	INT32_MIN := -(1 << 31)
	demos := []string{"2147483648", "-2147483648", "AB", " \n\t123B", " -1", "", "  ", "+1", "  0000000000012345678", "-001", "+-2", "-+1"}
	results := []int{INT32_MAX, INT32_MIN, 0, 123, -1, 0, 0, 1, 12345678, -1, 0, 0}
	for i, _ := range demos {
		result := myAtoi(demos[i])
		if result != results[i] {
			fmt.Printf("%d!=%d\n", result, results[i])
		} else {
			fmt.Printf("done %d=%d\n", result, results[i])
		}
	}
}

func myAtoi(str string) int {
	INT32_MAX := 1<<31 - 1
	INT32_MIN := -(1 << 31)
	index := 0
	strLen := len(str)
	temp := 0
	for value := 0; index < strLen && temp == 0; index++ {
		value = int(str[index])
		if value > 48 && value <= 57 { //1-9
			temp = value - 48
			break
		} else if value == 48 { //0
			temp = handleZero(&str, &index)
			if temp == -1 {
				return 0
			}
			break
		} else if value == 45 || value == 43 { //43是+， 44是-
			if index < strLen-1 && (str[index+1] >= 48 || str[index+1] <= 57) {
				temp = handleZero(&str, &index)
				if temp == -1 {
					return 0
				}
				if value == 45 {
					temp = -temp
				}
				break
			} else {
				return 0 //要么越界要么遇到非法字符
			}
		} else if value != 32 && (value < 9 || value > 11) {
			return 0
		}
	}
	if temp == 0 { //没有有效数字
		return 0
	}
	j := 1 //用于溢出检查
	value := 0
	for index = index + 1; index < strLen; index++ {
		value = int(str[index])
		if value < 48 || value > 57 {
			return temp
		}
		value -= 48
		if j > 8 {
			if temp > 0 {
				if temp > (INT32_MAX-value)/10 {
					return INT32_MAX
				}
			} else {
				if temp < (INT32_MIN+value)/10 {
					return INT32_MIN
				}
			}
		}
		if temp > 0 {
			temp = temp*10 + value
		} else {
			temp = temp*10 - value
		}
		j++
	}
	return temp
}

//注意是从下一个index开始比
func handleZero(str *string, index *int) int {
	var value int
	strLen := len(*str)
	var temp int
	for *index = *index + 1; *index < strLen; *index++ {
		value = int((*str)[*index])
		if value > 48 && value <= 57 {
			temp = value - 48
			break
		} else if value != 48 {
			return -1
		}
	}
	return temp
}
