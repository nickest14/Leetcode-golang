// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// ListNode is to define the struct
type ListNode = structures.ListNode

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cur := head

	for cur != nil && cur.Next != nil {
		node := &ListNode{Val: gcd(cur.Val, cur.Next.Val)}
		node.Next = cur.Next
		cur.Next = node
		cur = node.Next
	}
	return head
}
