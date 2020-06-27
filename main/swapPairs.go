package main

func swapPairs(head *ListNode) *ListNode {
	var help func(h *ListNode) *ListNode
	help = func(h *ListNode) *ListNode {
		if h != nil && h.Next != nil {
			temp := h.Next
			h.Next = help(h.Next.Next)
			temp.Next = h
			return temp
		} else {
			return h
		}
	}
	return help(head)
}
