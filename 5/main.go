package main

import (
	"fmt"
)

func main() {
	test()
}

func test() {
	demos := []string{"ababababa", "bb", "1abba", "abba1"}
	values := []string{"ababababa", "bb", "abba", "abba"}
	length := len(demos)
	var value string
	for i := 0; i < length; i++ {
		value = longestPalindrome(demos[i])
		if value == values[i] {
			fmt.Printf("equal: %s=%s \n", value, values[i])
		} else {
			fmt.Printf("not equal: %s!=%s\n", value, values[i])
		}
	}
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	strLen := len(s)
	buffer := make([]byte, 0, 2*strLen+3) //处理后的string
	buffer = append(buffer, '$', '#')
	for i := 0; i < strLen; i++ {
		buffer = append(buffer, s[i], '#')
	}
	buffer = append(buffer, '!')
	mx := 0
	id := 0 //字符串的index
	p := make([]int, 2*strLen+3, 2*strLen+3)
	var temp int
	var indexLeft int
	var indexRight int
	length := 2*strLen + 1
	for i := 1; i < length; i++ {
		if mx > i { //利用规律
			temp = p[2*id-i]
			if temp < mx-i { //注意不能是等，举个例子就知道了，
				//主要是利用对称的性质mx'q'...j...id...i...qmx。利用q和q'一定相等卡死了i继续往右扩张的可能性
				p[i] = temp //这就是最终的值了，此时mx和id很明显是没变的，直接continue
				continue
			} else {
				p[i] = mx - i
				indexRight = mx
				indexLeft = 2*i - indexRight
			}
		} else { //不能利用规律的，从旁边一个一个开始比
			p[i] = 1
			indexLeft = i - 1
			indexRight = i + 1
		}
		for buffer[indexLeft] == buffer[indexRight] {
			indexLeft--
			indexRight++
			// fmt.Printf("info-i%d\n", i)
			// fmt.Printf("info-left%d\n", indexLeft)
			// fmt.Printf("info-right%d\n", indexRight)
			p[i]++
		}
		if (indexRight - 1) > mx {
			mx = indexRight - 1
			id = i
		}
	}

	//现在检查p中的最大值
	valueMax := 0
	index := 0 //对应的p中的index
	for i := 1; i < length; i++ {
		if p[i] > valueMax {
			valueMax = p[i]
			index = i
		}
	}

	var startIndex int //字符串截取的开始
	if index%2 == 0 {
		startIndex = (index/2 - 1) - (valueMax/2 - 1)
	} else {
		startIndex = (index+1)/2 - 1 - (valueMax-1)/2
	}

	return s[startIndex : startIndex+valueMax-1] //实际的长度等于p中的value-1
}
