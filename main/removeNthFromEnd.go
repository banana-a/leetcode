package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head
	if head.Next == nil {
		return nil
	}
	for i := 0; i < n; i++ {
		q = q.Next
	}
	if q.Next == nil {
		return head.Next
	}
	for q.Next != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return head
}
