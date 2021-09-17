package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{1, 3, 2, 2, 3, 5, 2, 3, 5, 5, 5, 5, 5, 6, 7, 6, 7}
	//fmt.Println("I am lyming")
	m := SubSeq(nums)
	fmt.Println(m)
}

func SubSeq(nums []int) int {
	countMap := make(map[int]int)

	for _, v := range nums {
		countMap[v] += 1
	}

	max := 0

	for v, c := range countMap {
		mc := c + countMap[v+1]
		if mc > max {
			max = mc
		}
	}
	return max
}

func evalRPN(tokens []string) int64 {
	if len(tokens) <= 0 {
		return 0
	}
	stack := make([]int64, 0)

	for _, e := range tokens {
		switch e {
		case "+":
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack[len(stack)-2] = n1 + n2
			stack = stack[0 : len(stack)-1]
		case "-":
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack[len(stack)-2] = n1 - n2
			stack = stack[0 : len(stack)-1]
		case "*":
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack[len(stack)-2] = n1 * n2
			stack = stack[0 : len(stack)-1]
		case "/":
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack[len(stack)-2] = n1 / n2
			stack = stack[0 : len(stack)-1]
		default:
			v, err := strconv.ParseInt(e, 10, 64)
			if err != nil {
				panic(err.Error())
			}
			stack = append(stack, v)
		}
	}
	return stack[0]
}

func op(opToken string, n1, n2 int64) int64 {
	switch opToken {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	}
	panic("unexpected")
}
