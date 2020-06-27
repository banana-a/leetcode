package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var help func(h *ListNode) *ListNode
	help = func(h *ListNode) *ListNode {
		if h == nil {
			return nil
		}
		p := h
		for i := 1; i < k; i++ {
			if p.Next != nil {
				p = p.Next
			} else {
				return h
			}
		}
		temp := p.Next
		left := h
		right := left.Next
		for left != p {
			temp := right.Next
			right.Next = left
			left = right
			right = temp
		}
		h.Next = help(temp)
		return p
	}
	return help(head)
}
