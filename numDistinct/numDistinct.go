package main

import "fmt"

//115. 不同的子序列
//给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
//
//字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
//
//题目数据保证答案符合 32 位带符号整数范围。

/*
func numDistinct(s string, t string) int {
	return numDistinctHelper(s, 0, t, 0)
}

func numDistinctHelper(s string, startIndex int, t string, goalStartIndex int) int {
	if goalStartIndex == len(t) {
		return 1
	}

	if len(t)-goalStartIndex > len(s)-startIndex {
		return 0
	}

	num := 0
	for i := startIndex; i < len(s); i++ {
		if s[i] == t[goalStartIndex] {
			num += numDistinctHelper(s, i+1, t, goalStartIndex+1)
		}
	}
	return num
}
*/

func numDistinct(s string, t string) int {
	sLen := len(s)
	tLen := len(t)

	if sLen < tLen {
		return 0
	}

	dp := make([][]int, sLen+1, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]int, tLen+1, tLen+1)
	}

	for i := 0; i < sLen+1; i++ {
		dp[i][tLen] = 1
	}

	for j := 0; j < tLen; j++ {
		dp[sLen][j] = 0
	}

	for i := sLen - 1; i >= 0; i-- {
		for j := tLen - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j] + dp[i+1][j+1]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}

func main() {
	s := "rabbbit"
	t := "rabbit"
	x := numDistinct(s,t)
	fmt.Println(x)
}
