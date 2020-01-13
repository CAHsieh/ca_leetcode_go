package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, slow, fast := head, head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	prev.Next = nil

	left := sortList(head)
	right := sortList(slow)

	return mergeList(left, right)
}

func mergeList(left, right *ListNode) *ListNode {

	var head = new(ListNode)
	headTmp := head

	for left != nil && right != nil {
		if left.Val <= right.Val {
			head.Next = left
			left = left.Next
		} else {
			head.Next = right
			right = right.Next
		}
		head = head.Next
	}

	if left != nil {
		head.Next = left
	}

	if right != nil {
		head.Next = right
	}

	return headTmp.Next
}
