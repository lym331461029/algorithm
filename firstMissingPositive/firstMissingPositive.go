package main

import "fmt"

//41. 缺失的第一个正数
//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//https://leetcode-cn.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	if len(nums) <= 0 {
		return 1
	}

	num := len(nums)

	for i := 0; i < num; {
		if nums[i] > 0 && nums[i] <= num && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			x := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = x
			continue
		}
		i++
	}

	for i := 1; i <= num; i++ {
		if nums[i-1] != i {
			return i
		}
	}
	return num + 1
}
func main() {
	ints := []int{3, 4, -1, 1}
	x := firstMissingPositive(ints)
	fmt.Println(x)
}
