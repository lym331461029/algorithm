package main

import "fmt"

//135. 分发糖果
//老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//
//你需要按照以下要求，帮助老师给这些孩子分发糖果：
//
//每个孩子至少分配到 1 个糖果。
//评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
//
//那么这样下来，老师至少需要准备多少颗糖果呢？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/candy
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func candy(ratings []int) int {
	cNum := len(ratings)

	if cNum <= 0 {
		return 0
	}

	candies := make([]int, cNum, cNum)
	decIndex := 0
	candies[0] = 1
	for i := 1; i < cNum; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
			decIndex = i
		} else if ratings[i] == ratings[i-1] {
			candies[i] = 1
			decIndex = i
		} else {
			candies[i] = 1
			for j := i - 1; j >= decIndex; j-- {
				if candies[j] <= candies[j+1] {
					candies[j] += 1
				}
			}
		}
	}

	result := 0
	for _, v := range candies {
		result += v
	}
	fmt.Println(candies)
	return result
}

func main() {
	c := []int{1, 3, 2, 2, 1}
	x := candy(c)
	fmt.Println(x)
}
