package nonleetcode

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var operators []rune = []rune{'+', '-', '*', '/', '(', ')', '#'}

var operatorIndexMap map[rune]int

var priorityMatirx [][]int

const (
	PRI_X = iota // 错误的匹配，代表两个操作符不能进行优先级比价（不可能相遇）
	PRI_H
	PRI_L
	PRI_E
)

var ExpressionError error = errors.New("error expression")

func init() {
	// operatorIndexMap
	operatorIndexMap = map[rune]int{}
	for i := range operators {
		operatorIndexMap[operators[i]] = i
	}
	// priorityMatirx
	priorityMatirx = [][]int{
		// +	-	   *      /      (      )      #
		{PRI_H, PRI_H, PRI_L, PRI_L, PRI_L, PRI_H, PRI_H}, // +
		{PRI_H, PRI_H, PRI_L, PRI_L, PRI_L, PRI_H, PRI_H}, // -
		{PRI_H, PRI_H, PRI_H, PRI_H, PRI_L, PRI_H, PRI_H}, // *
		{PRI_H, PRI_H, PRI_H, PRI_H, PRI_L, PRI_H, PRI_H}, // /
		{PRI_L, PRI_L, PRI_L, PRI_L, PRI_L, PRI_E, PRI_X}, // (
		{PRI_H, PRI_H, PRI_H, PRI_H, PRI_X, PRI_E, PRI_H}, // )
		{PRI_L, PRI_L, PRI_L, PRI_L, PRI_L, PRI_X, PRI_E}, // #
	}
}

func mustGetOperatorPriority(one, two rune) int {
	idx1, ok1 := operatorIndexMap[one]
	idx2, ok2 := operatorIndexMap[two]
	if !ok1 || !ok2 {
		panic(fmt.Sprintf("not support operator: one:%c[ok=%v] two:%c[ok=%v]", idx1, ok1, idx2, ok2))
	}
	return priorityMatirx[idx1][idx2]
}

func isOperator(in rune) bool {
	_, ok := operatorIndexMap[in]
	return ok
}

// calTopNums2: 运算符栈顶出栈，nums出栈两个数字
func calTopNums2(operators *[]rune, nums *[]int) error {
	numLen := len(*nums)
	oprLen := len(*operators)
	num1 := (*nums)[numLen-1]
	num2 := (*nums)[numLen-2]
	*nums = (*nums)[:numLen-2]

	operator := (*operators)[oprLen-1]
	*operators = (*operators)[:oprLen-1]

	var numRes int
	switch operator { // not concern overflow
	case '+':
		numRes = num1 + num2
	case '-':
		numRes = num2 - num1
	case '*':
		numRes = num1 * num2
	case '/':
		numRes = num2 / num1
	default:
		return fmt.Errorf("unknown operator: %v", operator)
	}

	// 新num入栈
	*nums = append(*nums, numRes)

	return nil
}

func SimpleExpressionCal2(input string) (int, error) {
	// 消除表达式中的空格，在最后增加#标识
	input = strings.Replace(input, " ", "", -1)
	input = input + "#"
	if input[0] == '-' || input[0] == '+' {
		input = "0" + input
	}
	input = strings.Replace(input, "(-", "(0-", -1)
	input = strings.Replace(input, "(+", "(0+", -1)
	operators := []rune{'#'} // rune
	nums := []int{}          // int
	idx := 0
	inputLen := len(input)
	for idx < inputLen {
		val := rune(input[idx])
		currentTopOperator := operators[len(operators)-1]
		if isOperator(val) { // ok说明是运算符
			priority := mustGetOperatorPriority(currentTopOperator, val)
			switch priority {
			case PRI_H: // 操作符栈顶更高，计算结果
				calTopNums2(&operators, &nums)
				continue // idx不变
			case PRI_E: // 相等，区分#还是括号
				if currentTopOperator == '(' {
					operators = operators[:len(operators)-1]
					idx += 1
					continue
				}
				// 是#，就结束
				if currentTopOperator == '#' {
					if len(nums) == 0 {
						return 0, ExpressionError
					}
					return nums[0], nil
				}
			case PRI_L: // 栈顶优先级低，新运算符入栈
				operators = append(operators, val)
				currentTopOperator = val
			}
			idx += 1
			continue
		}
		// 说明这个byte是个数字，需要一直往后，得到整个数字
		startIndex := idx
		idx += 1
		for !isOperator(rune(input[idx])) {
			idx += 1
		}
		endIndex := idx
		numInt, err := strconv.Atoi(input[startIndex:endIndex])
		if err != nil {
			fmt.Printf("parse num error: %v", err)
			return 0, err
		}
		nums = append(nums, numInt)
	}
	return 0, errors.New("internal error")
}

// calTopNums: 运算符栈顶出栈，nums出栈两个数字
func calTopNums(operators, nums *list.List) error {
	numobj1 := nums.Front()
	if numobj1 == nil {
		return errors.New("num1 get failed")
	}
	num1 := numobj1.Value.(int)
	nums.Remove(numobj1)
	numobj2 := nums.Front()
	if numobj2 == nil {
		return errors.New("num2 get failed")
	}
	num2 := numobj2.Value.(int)
	nums.Remove(numobj2)

	operatorobj := operators.Front()
	if operatorobj == nil {
		return errors.New("operator get failed")
	}
	operator := operatorobj.Value.(rune)
	operators.Remove(operatorobj)

	var numRes int
	switch operator { // not concern overflow
	case '+':
		numRes = num1 + num2
	case '-':
		numRes = num2 - num1
	case '*':
		numRes = num1 * num2
	case '/':
		numRes = num2 / num1
	default:
		return fmt.Errorf("unknown operator: %v", operator)
	}

	// 新num入栈
	nums.PushFront(numRes)

	return nil
}

// SimpleExpressionCal 计算字符串表达式的值
// 处理负数和正数的两种情况：'-xxx'和'(-xxx'
func SimpleExpressionCal(input string) (int, error) {
	// 消除表达式中的空格，在最后增加#标识
	input = strings.Replace(input, " ", "", -1)
	input = input + "#"
	if input[0] == '-' || input[0] == '+' {
		input = "0" + input
	}
	input = strings.Replace(input, "(-", "(0-", -1)
	input = strings.Replace(input, "(+", "(0+", -1)
	operators := list.New() // rune
	nums := list.New()      // int
	operators.PushFront('#')
	idx := 0
	inputLen := len(input)
	for idx < inputLen {
		val := rune(input[idx])
		currentTopOperator := operators.Front().Value.(rune)
		if isOperator(val) { // ok说明是运算符
			priority := mustGetOperatorPriority(currentTopOperator, val)
			switch priority {
			case PRI_H: // 操作符栈顶更高，计算结果
				calTopNums(operators, nums)
				continue // idx不变
			case PRI_E: // 相等，区分#还是括号
				if currentTopOperator == '(' {
					top := operators.Front()
					operators.Remove((top))
					idx += 1
					continue
				}
				// 是#，就结束
				if currentTopOperator == '#' {
					numobj := nums.Front()
					if numobj == nil {
						return 0, ExpressionError
					}
					return numobj.Value.(int), nil
				}
			case PRI_L: // 栈顶优先级低，新运算符入栈
				operators.PushFront(val)
				currentTopOperator = rune(val)
			}
			idx += 1
			continue
		}
		// 说明这个byte是个数字，需要一直往后，得到整个数字
		startIndex := idx
		idx += 1
		for !isOperator(rune(input[idx])) {
			idx += 1
		}
		endIndex := idx
		numInt, err := strconv.Atoi(input[startIndex:endIndex])
		if err != nil {
			fmt.Printf("parse num error: %v", err)
			return 0, err
		}
		nums.PushFront(numInt)
	}
	return 0, errors.New("internal error")
}
