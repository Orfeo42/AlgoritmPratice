package leetcode

// https://leetcode.com/problems/swap-nodes-in-pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var prev *ListNode

	item := head
	head = item.Next
	for item != nil && item.Next != nil {
		tmp := item.Next
		if prev != nil {
			prev.Next = tmp
		}
		prev = item
		item.Next = item.Next.Next
		tmp.Next = item
		item = item.Next
	}
	return head
}
