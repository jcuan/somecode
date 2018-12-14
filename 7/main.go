package main

import (
	"fmt"
)

func main() {
	test()
}

func test() {
	demos := []int{0, 1, -1, 12, -12, 123, -123, 123456, 2147483647, -2147483647, -1563847412}
	results := []int{0, 1, -1, 21, -21, 321, -321, 654321, 0, 0, 0}
	for i, _ := range demos {
		result := reverse(demos[i])
		if result != results[i] {
			fmt.Printf("%d!=%d\n", result, results[i])
		} else {
			fmt.Printf("done %d=%d\n", result, results[i])
		}
	}
}

func reverse(x int) int {
	INT32_MAX := 1<<31 - 1
	INT32_MIN := -(1 << 31)
	temp := x / 10
	result := x % 10
	for result == 0 && temp != 0 {
		result = temp % 10
		temp /= 10
	}
	remainder := 0
	for i := 1; temp != 0; i++ {
		remainder = temp % 10
		if i > 8 {
			if result > 0 {
				if result > (INT32_MAX-remainder)/10 {
					return 0
				}
			} else {
				if result < (INT32_MIN-remainder)/10 {
					return 0
				}
			}
		}
		result = result*10 + remainder
		temp /= 10
	}
	return result
}
