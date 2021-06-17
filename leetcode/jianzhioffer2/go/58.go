package jzoffer

// https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/

func ReverseWords2(s string) string {
	buffer := make([]byte, len(s)+1) // 有可能多存一个空格
	bufferIndex := 0
	blank := byte(' ')
	index := len(s) - 1
	for index >= 0 {
		if s[index] == blank {
			index -= 1
			continue
		}
		endIndex := index
		for s[index] != blank {
			index -= 1
			if index == -1 {
				break
			}
		}
		for i := index + 1; i <= endIndex; i++ {
			buffer[bufferIndex] = s[i]
			bufferIndex += 1
		}
		// 每一个单词最后都有空格
		buffer[bufferIndex] = blank
		bufferIndex += 1
	}
	if bufferIndex <= 0 {
		return ""
	}
	return string(buffer[:bufferIndex-1]) // 最后多了一个空格
}

// 反有可能是方向的反
func ReverseWords(s string) string {
	currentWord := []byte{}
	savedWords := [][]byte{}
	blank := byte(' ')
	for i := range s {
		if s[i] == blank {
			if len(currentWord) != 0 {
				savedWords = append(savedWords, currentWord)
				currentWord = []byte{}
			}
		} else {
			currentWord = append(currentWord, s[i])
		}
	}
	if len(currentWord) != 0 {
		savedWords = append(savedWords, currentWord)
	}
	finalWords := []byte{}
	indexMax := len(savedWords)
	for i := indexMax - 1; i >= 0; i-- {
		savedWords[i] = append(savedWords[i], byte(' '))
		finalWords = append(finalWords, savedWords[i]...)
	}
	if indexMax > 0 {
		finalWords = finalWords[:len(finalWords)-1]
	}
	return string(finalWords)
}
