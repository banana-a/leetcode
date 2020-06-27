package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	mergeTwo := func(l1 *ListNode, l2 *ListNode) *ListNode {
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
		if l1 != nil {
			p.Next = l1
		}
		if l2 != nil {
			p.Next = l2
		}
		return res.Next
	}
	for i := 1; i < len(lists); i++ {
		lists[0] = mergeTwo(lists[0], lists[i])
	}
	return lists[0]
}
