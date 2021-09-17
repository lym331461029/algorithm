package main

import (
	"fmt"
)

//164. 最大间距
//
//给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
//
//如果数组元素个数小于 2，则返回 0。

//二进制基数排序
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sortedNumGroups := [2][]int{}
	for i := 0; i < 32; i++ {
		var x int = 1 << i

		for _, num := range nums {
			if num&x == 0 {
				sortedNumGroups[0] = append(sortedNumGroups[0],num)
			} else {
				sortedNumGroups[1] = append(sortedNumGroups[1],num)
			}
		}

		l := copy(nums,sortedNumGroups[0])
		copy(nums[l:],sortedNumGroups[1])


		sortedNumGroups[0] = sortedNumGroups[0][0:0]
		sortedNumGroups[1] = sortedNumGroups[1][0:0]
	}

	x := 0
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i+1] - nums[i] > x {
			x = nums[i+1] - nums[i]
		}
	}
	return x
}

func main()  {
	var x []int = []int{1,1,1,1}
	fmt.Println(x)
	r := maximumGap(x)
	fmt.Println(x,r)
}
