package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p := res
	for l1 != nil && l2 != nil {
		var item *ListNode
		if l1.Val < l2.Val {
			item = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			item = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		p.Next = item
		p = p.Next
	}
	for l1 != nil {
		item := &ListNode{Val: l1.Val}
		p.Next = item
		p = p.Next
		l1 = l1.Next
	}
	for l2 != nil {
		item := &ListNode{Val: l2.Val}
		p.Next = item
		p = p.Next
		l2 = l2.Next
	}
	return res.Next
}
