package main

//79. 单词搜索 leetcode
//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/word-search
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func exist(board [][]byte, word string) bool {

	row := len(board)
	line := len(board[0])

	flag := make([][]int8, row, row)
	for i := 0; i < row; i++ {
		flag[i] = make([]int8, line, line)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if match(board, i, j, []byte(word), flag) {
				return true
			}
		}
	}
	return false
}

func match(board [][]byte, startRow int, startLine int, goal []byte, flag [][]int8) bool {
	if len(goal) <= 0 {
		return true
	}
	row := len(board)
	line := len(board[0])

	if startRow < 0 || startRow >= row ||
		startLine < 0 || startLine >= line ||
		board[startRow][startLine] != goal[0] ||
		flag[startRow][startLine] > 0 {
		return false
	}

	flag[startRow][startLine] = 1
	if match(board, startRow+1, startLine, goal[1:], flag) ||
		match(board, startRow-1, startLine, goal[1:], flag) ||
		match(board, startRow, startLine+1, goal[1:], flag) ||
		match(board, startRow, startLine-1, goal[1:], flag) {
		return true
	} else {
		flag[startRow][startLine] = 0
		return false
	}
}

func main() {
	//
}
