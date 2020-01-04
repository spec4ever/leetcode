package top

/*
No: 206
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre, after *ListNode

	for head != nil {
		after = head.Next
		head.Next = pre
		pre = head
		head = after
	}

	return pre

}
