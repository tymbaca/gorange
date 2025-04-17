package main

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	str := "[" + strconv.Itoa(n.Val)
	cur := n.Next

	for cur != nil {
		str += ", " + strconv.Itoa(cur.Val)
		cur = cur.Next
	}
	str += "]"

	return str
}

func fromSlice(nums []int) *ListNode {
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}

	prev := head
	for i, n := range nums {
		if i == 0 {
			continue
		}

		prev.Next = &ListNode{
			Val: n,
		}

		prev = prev.Next
	}

	return head
}
