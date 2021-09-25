package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type arrListNode []*ListNode

func (arr arrListNode) Len() int {
	return len(arr)
}

func (arr arrListNode) Less(i, j int) bool {
	return arr[i].Val < arr[j].Val
}

func (arr arrListNode) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr arrListNode) Push(x interface{}) {
	arr = append(arr, x.(*ListNode))
}

func (arr arrListNode) Pop() interface{} {
	x := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return x
}

func (arr arrListNode) liitleHead() {
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i].Val > arr[index].Val {
			index = i
		}
	}
	arr[0], arr[index] = arr[index], arr[0]
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := ListNode{}
	tail := &head

	if len(lists) <= 0 {
		return head.Next
	}

	arrList := arrListNode{}
	for _, l := range lists {
		if l != nil {
			arrList = append(arrList, l)
		}
	}
	heap.Init(arrList)

	for len(arrList) > 0 {
		min := arrList[0]
		if min.Next != nil {
			arrList[0] = min.Next
		} else {
			arrList = arrList[1:]
		}
		heap.Init(arrList)
		tail.Next = min
		tail = tail.Next
	}
	tail.Next = nil
	return head.Next
}

func main() {
	lists := []*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	result := mergeKLists(lists)
	printList(result)

}
func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%v -> ", l.Val)
		l = l.Next
	}
}
