package linkedlist

// DeleteNode deletes node without pointer to the Head
func DeleteNode(node *ListNode) {
	// 1 solution
	node.Val = node.Next.Val
	node.Next = node.Next.Next

	// 2 solution
	*(node) = *(node.Next)
}