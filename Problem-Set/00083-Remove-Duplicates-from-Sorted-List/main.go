package main

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	hash := map[int]bool{}
	res := &ListNode{}
	current := res

	for head != nil {
		val := head.Val

		if _, ok := hash[val]; !ok {
			hash[val] = true
			current.Next = &ListNode{Val: val}
			current = current.Next
		}

		head = head.Next
	}

	return res.Next
}
