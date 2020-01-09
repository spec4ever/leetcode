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
思路： 递归，过程+终止
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

/*
No: 21
思路： 迭代，头结点便于返回
*/
func mergeTwoListsIter(l1 *ListNode, l2 *ListNode) *ListNode {

	var head, cur *ListNode

	head = &ListNode{Val: -1, Next: nil}
	cur = head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}

	return head.Next
}

/*
No: 88 合并两个有序数组
*/
/*
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	mPointer := m
	nPointer := n
	for mPointer > 0 && nPointer > 0 {
		if nums1[mPointer-1] < nums2[nPointer-1] {
			nums1[mPointer+nPointer-1] = nums2[nPointer-1]
			nPointer--
		} else {
			nums1[mPointer+nPointer-1] = nums1[mPointer-1]
			mPointer--
		}
		if mPointer == 0 && nPointer > 0 {
			for nPointer > 0 {
				nums1[mPointer+nPointer-1] = nums2[nPointer-1]
				nPointer--
			}
		}
	}
}

/*
No: 202 快乐数
思路： 快乐数最后是1循环，非快乐数也会基于某个数循环，快慢指针破循环，好处是无须用额外存储空间
*/
func isHappy(n int) bool {

	slow := bitSquareSum(n)
	fast := bitSquareSum(bitSquareSum(n))
	for slow != fast {
		slow = bitSquareSum(slow)
		fast = bitSquareSum(bitSquareSum(fast))
	}

	if slow == 1 {
		return true
	}

	return false

}
func bitSquareSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}

	return sum
}
