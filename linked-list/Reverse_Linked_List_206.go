package linkedlist

// ReverseList ...
func ReverseList(head *ListNode) *ListNode {
	// Iterative Solution: Space O(1), Time O(n)
	/*
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	prev := &ListNode{}

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
	*/

	// Recursive Solution: Space O(n), Time O(n)
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseList(head.Next)

	head.Next.Next = head
	head.Next = nil
	return newHead
}