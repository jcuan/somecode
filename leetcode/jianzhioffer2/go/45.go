package jzoffer

import (
	"bytes"
	"sort"
	"strconv"
)

type ValDetail struct {
	Val   int // 实际的值
	Digit int // 位数
}

type MySlice []ValDetail

func (this MySlice) Len() int {
	return len(this)
}

func (this MySlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// 计算base*10^zhishu大小
func pow(base, zhishu int) int64 {
	multi := 1
	for ; zhishu > 0; zhishu-- {
		multi *= 10
	}
	return int64(base) * int64(multi)
}

func (this MySlice) Less(i, j int) bool {
	return pow(this[i].Val, this[j].Digit)+int64(this[j].Val) < pow(this[j].Val, this[i].Digit)+int64(this[i].Val)
}

// https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/solution/mian-shi-ti-45-ba-shu-zu-pai-cheng-zui-xiao-de-s-4/
func minNumber(nums []int) string {
	items := make([]ValDetail, len(nums))
	for i, val := range nums { // 将每一个的每一位提取出来
		valDetail := ValDetail{val, 0} // 0的dist也应该为1
		for val > 0 {
			valDetail.Digit++
			val /= 10
		}
		if valDetail.Digit == 0 {
			valDetail.Digit = 1
		}
		items[i] = valDetail
	}
	myslice := MySlice(items)
	sort.Sort(myslice)
	var buf bytes.Buffer
	for _, val := range items {
		buf.WriteString(strconv.Itoa(val.Val))
	}
	return buf.String()
}
