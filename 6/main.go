package main

import (
	"fmt"
)

func main() {
	test()
}

func test() {
	demos := []string{"ABCDEFGHIJKLMN", "LEETCODEISHIRING", "LEETCODEISHIRING", "A"}
	rows := []int{4, 3, 4, 1}
	results := []string{"AGMBFHLNCEIKDJ", "LCIRETOESIIGEDHN", "LDREOEIIECIHNTSG", "A"}
	for index, _ := range demos {
		result := convert(demos[index], rows[index])
		if result == results[index] {
			fmt.Printf("%s=%s\n", result, results[index])
		} else {
			fmt.Printf("%s!=%s\n", result, results[index])
		}
	}
}

// A    I
// B   H
// C  G
// D F
// E
//....................
// A - H这种叫一个circle
// main：A-E这种列
func convert(s string, numRows int) string {
	strLen := len(s)
	if strLen < 3 || numRows == 1 {
		return s
	}
	circleLen := 2*numRows - 2 //2r-2是一个完整的|/的长度
	buffer := make([]byte, strLen, strLen)
	bufferIndex := 0
	temp := 0
	for i := 0; i < numRows; i++ {
		j := 0
		if i != 0 && i != numRows-1 {
			mainTurn := true //应该读取main
			mainEnd := false
			singleEnd := false
			temp = i
			for {
				if mainTurn {
					mainTurn = false
					if mainEnd == false {
						temp = j*circleLen + i
						if temp < strLen {
							buffer[bufferIndex] = s[temp]
							bufferIndex++
						} else {
							mainEnd = true
						}
					}
				} else {
					mainTurn = true
					if singleEnd == false {
						temp = (j+1)*circleLen - i
						j++
						if temp < strLen {
							buffer[bufferIndex] = s[temp]
							bufferIndex++
						} else {
							singleEnd = true
						}
					}
				}
				if mainEnd && singleEnd {
					break
				}
			}
		} else {
			temp = i
			for temp < strLen {
				buffer[bufferIndex] = byte(s[temp])
				bufferIndex++
				temp += circleLen
			}
		}

	}
	return string(buffer)
}

// 感觉这个解法自己真是智障了
// A    I
// B   H
// C  G
// D F
// E
//mainColumn：A-B这种列
//singleColumn：F、G、H所在的列
func convertBack(s string, numRows int) string {
	strLen := len(s)
	if strLen < 3 || numRows == 1 {
		return s
	}
	var mainColumnCount, singleColumnCount int
	var mainEleCount int                       //mainColumn中所有元素的个数
	mainColumnCount = strLen / (2*numRows - 2) //2r-2是一个完整的|/的
	remainder := strLen % (2*numRows - 2)
	if remainder > numRows {
		singleColumnCount = mainColumnCount*(numRows-2) + (remainder - numRows)
		mainEleCount = mainColumnCount*numRows + numRows
	} else {
		mainEleCount = mainColumnCount*numRows + remainder
		singleColumnCount = mainColumnCount * (numRows - 2)
	}
	if remainder > 0 {
		mainColumnCount++
	}

	mainList := make([]byte, mainEleCount, mainEleCount)
	singleList := make([]byte, singleColumnCount, singleColumnCount)

	mainIndex := 0
	singleIndex := 0
	indicator := 0 //实在不知道该叫啥
	num := 2*numRows - 2
	for _, ele := range s {
		if indicator < numRows {
			mainList[mainIndex] = byte(ele)
			mainIndex++
		} else {
			singleList[singleIndex] = byte(ele)
			singleIndex++
		}
		indicator = (indicator + 1) % num
	}

	//现在开始读出来
	buffer := make([]byte, strLen, strLen)
	bufferIndex := 0
	temp := 0 //用来保存当前访问到的位置
	for i := 0; i < numRows; i++ {
		j := -1
		mainTurn := true              //该读mainList
		if i != 0 && i != numRows-1 { //从singleList中读取
			singleEnd := false //只是是否读取到这一行的最后，在读取就要越界了
			mainEnd := false
			for {
				if mainTurn { //从mainlist读取
					mainTurn = false
					j++
					if mainEnd == false {
						temp = j*numRows + i
						if temp < mainEleCount {
							buffer[bufferIndex] = mainList[temp]
							bufferIndex++
						} else {
							mainEnd = true
						}
					}
				} else { //该从singleList中读取
					mainTurn = true
					if singleEnd == false {
						temp = j*(numRows-2) + numRows - 2 - i
						if temp < singleColumnCount {
							buffer[bufferIndex] = singleList[temp]
							bufferIndex++
						} else {
							singleEnd = true
						}
					}
				}
				if singleEnd && mainEnd {
					break
				}
			}
		} else {
			temp = i
			for temp < mainEleCount {
				buffer[bufferIndex] = mainList[temp]
				bufferIndex++
				temp += numRows
			}
		}
	}
	return string(buffer)
}
