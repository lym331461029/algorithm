package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	index := findPeakElement(nums)
	fmt.Println(index)
}

func findPeakElement(nums []int) int {
	l := len(nums)
	if l <= 0 {
		return -1
	}

	if l == 1 {
		return 0
	}

	middle := l/2 - 1
	if nums[middle] < nums[middle+1] {
		return middle + 1 + findPeakElement(nums[middle+1:])
	} else {
		return findPeakElement(nums[:middle+1])
	}
}
