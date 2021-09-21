package main

import "fmt"

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
