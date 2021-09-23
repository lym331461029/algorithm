package main

//127. 单词接龙
//字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列：
//
//序列中第一个单词是 beginWord 。
//序列中最后一个单词是 endWord 。
//每次转换只能改变一个字母。
//转换过程中的中间单词必须是字典 wordList 中的单词。
//
//给你两个单词 beginWord 和 endWord 和一个字典 wordList ，找到从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/word-ladder
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordIDMap := map[string]int{}
	edge := [][]int{}

	nodeNum := 0

	for _, word := range wordList {
		wordIDMap[word] = nodeNum
		edge = append(edge, []int{})

		wordID := nodeNum
		nodeNum++

		wLen := len(word)
		for i := 0; i < wLen; i++ {
			wWord := word[:i] + "*" + word[i+1:]

			if wWordID, e := wordIDMap[wWord]; !e {
				wordIDMap[wWord] = nodeNum
				edge = append(edge, []int{})

				edge[wordID] = append(edge[wordID], nodeNum)
				edge[nodeNum] = append(edge[nodeNum], wordID)

				nodeNum++
			} else {
				edge[wordID] = append(edge[wordID], wWordID)
				edge[wWordID] = append(edge[wWordID], wordID)
			}

		}
	}

	wordIDMap[beginWord] = nodeNum
	edge = append(edge, []int{})
	beginWordID := nodeNum
	nodeNum++

	for i := 0; i < len(beginWord); i++ {
		wWord := beginWord[:i] + "*" + beginWord[i+1:]

		if wWordID, e := wordIDMap[wWord]; !e {
			wordIDMap[wWord] = nodeNum
			edge = append(edge, []int{})

			edge[beginWordID] = append(edge[beginWordID], nodeNum)
			edge[nodeNum] = append(edge[nodeNum], beginWordID)

			nodeNum++
		} else {
			edge[beginWordID] = append(edge[beginWordID], wWordID)
			edge[wWordID] = append(edge[wWordID], beginWordID)
		}
	}

	endWordId, e := wordIDMap[endWord]
	if !e {
		return 0
	}

	queue := []int{beginWordID}
	dis := make([]int, nodeNum, nodeNum)
	for i := 0; i < nodeNum; i++ {
		dis[i] = -1
	}
	dis[beginWordID] = 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == endWordId {
			return dis[endWordId]/2 + 1
		} else {
			for _, id := range edge[cur] {
				if dis[id] == -1 {
					dis[id] += dis[cur] + 1
					queue = append(queue, id)
				}
			}
		}
	}
	return 0
}
