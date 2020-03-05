package main

func main() {
	e := ListNode{5, nil}
	d := ListNode{4, &e}
	c := ListNode{3, &d}
	b := ListNode{2, &c}
	a := ListNode{1, &b}
	n := oddEvenList(&a)
	for ; n != nil; n = n.Next {
		println(n.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	firstEven := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next // 1 -> 3 -> 4
		odd = odd.Next       // 3 -> 4
		even.Next = odd.Next // 2 -> 4
		even = even.Next     // 4 -> 5
	}
	odd.Next = firstEven
	return head
}
