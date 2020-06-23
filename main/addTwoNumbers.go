package main

type ListNode struct {
	Next *ListNode
	Val int
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := &ListNode{
		Next: nil,
		Val:  0,
	}
	res := p
	up := 0
	for l1 != nil && l2 != nil {
		temp := (l1.Val + l2.Val + up) / 10
		num := (l1.Val + l2.Val + up) % 10
		item := &ListNode{
			Next: nil,
			Val:  num,
		}
		up = temp
		p.Next = item
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		temp := (l1.Val + up) / 10
		num := (l1.Val + up) % 10
		item := &ListNode{
			Next: nil,
			Val:  num,
		}
		up = temp
		p.Next = item
		p = p.Next
		l1 = l1.Next
	}
	for l2 != nil {
		temp := (l2.Val + up) / 10
		num := (l2.Val + up) % 10
		item := &ListNode{
			Next: nil,
			Val:  num,
		}
		up = temp
		p.Next = item
		p = p.Next
		l2 = l2.Next
	}
	if up == 1 {
		item := &ListNode{
			Next: nil,
			Val:  1,
		}
		p.Next = item
	}
	return res.Next
}