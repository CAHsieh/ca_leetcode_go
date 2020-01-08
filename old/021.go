package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 && nil == l2 {
		return nil
	} else if nil == l1 {
		return l2
	} else if nil == l2 {
		return l1
	}
	head := new(ListNode)
	tmp := new(ListNode)

	for nil != l1 && nil != l2 {
		if l1.Val <= l2.Val {
			if nil == head.Next {
				head.Next = l1
				tmp = l1
			} else {
				tmp.Next = l1
				tmp = l1
			}
			l1 = l1.Next
		} else {
			if nil == head.Next {
				head.Next = l2
				tmp = l2
			} else {
				tmp.Next = l2
				tmp = l2
			}
			l2 = l2.Next
		}
	}

	if nil != l1 {
		tmp.Next = l1
	} else if nil != l2 {
		tmp.Next = l2
	}

	return head.Next
}
