package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)

	a := nums[0] - 1

	var result [][]int
	for i := 0; i < len(nums)-2; {
		for nums[i] == a && i < len(nums)-2{
			i++
		}
		a = nums[i]

		y := 0 - a

		for j, k := i+1, len(nums)-1; j < k && nums[j] <= y; {
			if nums[j]+nums[k] == y {
				result = append(result, []int{a, nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if nums[j]+nums[k] > y {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[j]+nums[k] < y {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return result
}

func main() {
	a := []int{-4, -1, -1, 0, 1, 2}
	//a := []int{0, 0, 0}
	result := threeSum(a)
	fmt.Println(result)
}
