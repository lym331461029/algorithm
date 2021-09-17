package main

import "fmt"

func removeDuplicate(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	slowIndex := 0
	fastIndex := 0

	preElem := nums[fastIndex] - 1
	for {
		for fastIndex < len(nums) && nums[fastIndex] == preElem {
			fastIndex++
		}

		if fastIndex == len(nums) {
			break
		}
		if slowIndex != fastIndex {
			nums[slowIndex] = nums[fastIndex]
		}

		preElem = nums[fastIndex]
		slowIndex++
	}

	return slowIndex
}

func main() {
	a := []int{1, 2, 3, 3, 4, 5, 6, 6, 6, 7}
	fmt.Println(a)
	n := removeDuplicate(a)
	fmt.Println(a)
	fmt.Println(n)
	fmt.Println("Single Elem")
}
