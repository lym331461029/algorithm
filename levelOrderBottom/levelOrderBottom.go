package main


//107. 二叉树的层序遍历 II
//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	leveldElem := [][]int{}
	queue := []*TreeNode{
		root,
	}

	for len(queue) > 0 {
		size := len(queue)
		level := []int{}
		for i := 0; i < size; i++ {
			level = append(level, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		leveldElem = append(leveldElem, level)
	}

	levelCount := len(leveldElem)
	for i, j := 0, levelCount-1; i < j; {
		leveldElem[i], leveldElem[j] = leveldElem[j], leveldElem[i]
		i++
		j--
	}
	return leveldElem
}
