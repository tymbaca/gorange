package main

import (
	"fmt"
	"strconv"
)

func main() {
	list := fromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Printf("%s\n", list)

	reversed := reverseList(list)
	fmt.Printf("%s\n", reversed)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, middle, right := head, head.Next, head.Next.Next
	left.Next = nil
	for right != nil {
		fmt.Println("loop")
		middle.Next = left

		left = middle
		middle = right
		right = right.Next
	}
	middle.Next = left

	return middle
}

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
