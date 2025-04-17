package main

func main() {
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	prev := list1
	merge := list2

	if prev.Val > merge.Val {
		prev, merge = merge, prev
	}

	head := prev.Next
	nextMerge := merge.Next

	for {
		if merge.Val < head.Val {
			prev.Next = merge
			merge.Next = head

			head = head.Next
			prev = merge
		} else {
			merge.Next = head.Next
			head.Next = merge

			prev = head
			head = merge
		}

		merge = nextMerge
		nextMerge = merge.Next
	}
}
