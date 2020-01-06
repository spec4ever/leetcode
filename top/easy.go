package top

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
No: 1
*/
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	hashMap := map[int]int{}
	for k, v := range nums {
		if index, ok := hashMap[v]; ok {
			return []int{index, k}
		} else {
			hashMap[target-v] = k
		}
	}

	return nil
}

/*
No: 206
*/
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

/*
No: 21
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	//终止
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
