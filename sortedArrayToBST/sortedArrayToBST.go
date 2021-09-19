package main

import (
	"fmt"
)

//108. 将有序数组转换为二叉搜索树
//给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
//
//高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)

	if l == 0 {
		return nil
	} else if l == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	} else {
		middle := l / 2
		return &TreeNode{
			Val:   nums[middle],
			Left:  sortedArrayToBST(nums[:middle]),
			Right: sortedArrayToBST(nums[middle+1:]),
		}
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}


//109. 有序链表转换二叉搜索树
func sortedListToBST2(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
		}
	} else if head.Next.Next == nil {
		return &TreeNode{
			Val: head.Next.Val,
			Left: &TreeNode{
				Val: head.Val,
			},
		}
	} else {
		t1, t2 := head, head.Next.Next
		for t2 != nil && t2.Next != nil {
			t1 = t1.Next
			t2 = t2.Next
			if t2 != nil {
				t2 = t2.Next
			} else {
				break
			}
		}

		node := &TreeNode{
			Val: t1.Next.Val,
		}
		right := t1.Next.Next
		t1.Next = nil
		node.Left = sortedListToBST2(head)
		node.Right = sortedListToBST2(right)
		return node
	}
}

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			//Next: &ListNode{
			//	Val: 3,
			//},
		},
	}

	t := sortedListToBST2(h)
	fmt.Println(t)
}
