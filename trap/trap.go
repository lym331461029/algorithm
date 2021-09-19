package main

import "fmt"

//42. 接雨水
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

func trap(height []int) int {
	length := len(height)
	if length < 3 {
		return 0
	}

	highestIndex := 0

	traps := make([]int, length, length)

	for i := 1; i < length; i++ {
		if height[i] > height[i-1] {

			nearestHigher := i
			for n := i - 2; n >= 0; n-- {
				if height[n] >= height[i] {
					nearestHigher = n
					break
				}
			}

			if nearestHigher == i {
				nearestHigher = highestIndex
			}
			level := min(height[nearestHigher], height[i])

			for u := nearestHigher; u < i; u++ {
				traps[u] = max(0, level-height[u])
			}
		}

		if height[i] >= height[highestIndex] {
			highestIndex = i
		}
	}

	result := 0

	for _, v := range traps {
		result += v
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	height := []int{4,3,3,9,3,0,9,2,8,3}
	x := trap(height)
	fmt.Println(x)
}
