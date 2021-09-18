package main

import "fmt"

//174. 地下城游戏
//一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N 个房间组成的二维网格。我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。
//
//骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。
//
//有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
//
//为了尽快到达公主，骑士决定每次只向右或向下移动一步。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/dungeon-game
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) <= 0 || len(dungeon[0]) <= 0 {
		return 0
	}

	row := len(dungeon)
	line := len(dungeon[0])
	return calculateMinimumHPHelper(dungeon, row-1, line-1, 1)
}

func calculateMinimumHPHelper(dungeon [][]int, goalRow int, goalLine int, needSorece int) int {
	if goalRow < 0 || goalRow >= len(dungeon) {
		return 0
	}

	if goalLine < 0 || goalLine >= len(dungeon[0]) {

		return 0
	}

	needSorece = needSorece - dungeon[goalRow][goalLine]

	if needSorece <= 1 {
		needSorece = 1
	}

	upNeedScore := calculateMinimumHPHelper(dungeon, goalRow-1, goalLine, needSorece)

	leftNeddScore := calculateMinimumHPHelper(dungeon, goalRow, goalLine-1, needSorece)

	if upNeedScore == 0 && leftNeddScore == 0 {
		return needSorece
	} else if upNeedScore != 0 && leftNeddScore == 0 {
		return upNeedScore
	} else if upNeedScore == 0 && leftNeddScore != 0 {
		return leftNeddScore
	} else {
		if upNeedScore < leftNeddScore {
			return upNeedScore
		} else {
			return leftNeddScore
		}
	}
}

func calculateMinimumHP2(dungeon [][]int) int {
	if len(dungeon) <= 0 || len(dungeon[0]) <= 0 {
		return 0
	}

	row := len(dungeon)
	line := len(dungeon[0])

	dp := make([][]int, row, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, line, line)
	}

	for i := row - 1; i >= 0; i-- {
		for j := line - 1; j >= 0; j-- {
			if i == row-1 && j == line-1 {
				dp[i][j] = max(1, 1-dungeon[i][j])
			} else if i == row-1 {
				dp[i][j] = max(1, dp[i][j+1]-dungeon[i][j])
			} else if j == line-1 {
				dp[i][j] = max(1, dp[i+1][j]-dungeon[i][j])
			} else {
				dp[i][j] = max(1,min(dp[i][j+1]-dungeon[i][j],dp[i+1][j]-dungeon[i][j]))
			}
		}
	}
	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
//[[-2,-3,3],[-5,-10,1],[10,30,-5]]
func main() {
	dungeon := [][]int{
		[]int{-2, -3, 3},
		[]int{-5, -10, 1},
		[]int{10, 30, -5},
	}

	//dungeon2 := [][]int{
	//	[]int{-2, -3},
	//	[]int{10, -5},
	//}
	x := calculateMinimumHP2(dungeon)
	fmt.Println(x)
}
