package main

import (
	"fmt"
)

func main() {
	//h := ConstructList([]int{1, 2, 3, 4})
	//PrintList(h)
	//
	//h1, _ := reverseListHelper(h)
	//PrintList(h1)

	h := ConstructList([]int{1, 2, 8, 9, 10, 15, 5, 3, 4})
	PrintList(h)

	h1 := partition(h, 5)
	PrintList(h1)

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Next *ListNode
	Val  int
}

func ReverseList(head *ListNode) *ListNode {
	h, _ := reverseListHelper(head)
	return h
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("=> %d ", head.Val)
		head = head.Next
	}
	fmt.Println("")
}

func ConstructList(els []int) *ListNode {
	var head, cur *ListNode

	for _, v := range els {
		if head == nil {
			head = &ListNode{
				Val: v,
			}
			cur = head
		} else {
			cur.Next = &ListNode{
				Val: v,
			}
			cur = cur.Next
		}
	}
	return head
}

func reverseListHelper(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	if head.Next != nil {
		h, last := reverseListHelper(head.Next)
		last.Next = head
		head.Next = nil
		return h, last.Next
	} else {
		return head, head
	}
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lessHead := &ListNode{}
	lastLess := lessHead
	greaterHead := &ListNode{}
	lastGreater := greaterHead

	cur := head
	for cur != nil {
		if cur.Val < x {
			lastLess.Next = cur
			lastLess = lastLess.Next
		} else {
			lastGreater.Next = cur
			lastGreater = lastGreater.Next
		}
		cur = cur.Next
	}
	lastGreater.Next = nil
	lastLess.Next = nil

	if lastGreater != greaterHead {
		lastLess.Next = greaterHead.Next
	}

	return lessHead.Next
}
