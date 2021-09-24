package main

import "fmt"

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)

	dp := make([][]bool, sLen+1, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]bool, pLen+1, pLen+1)
	}

	dp[0][0] = true
	for j := 1; j < pLen+1; j++ {
		dp[0][j] = dp[0][j-1] && p[j-1] == '*'
	}

	for i := 1; i < sLen+1; i++ {
		for j := 1; j < pLen+1; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-1] && (p[j-1] == '?' || p[j-1] == s[i-1])
			}
		}
	}
	return dp[sLen][pLen]
}

func main() {
	s := "cb"
	p := "?a"
	x := isMatch(s, p)
	fmt.Println(x)
}
