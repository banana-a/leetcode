package main

func detectCycle(head *ListNode) *ListNode {
	p, q := head, head
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
		if p == q {
			p = head
			break
		}
	}
	if q == nil || q.Next == nil {
		return nil
	}
	for p != q {
		p = p.Next
		q = q.Next
	}
	return p
}
