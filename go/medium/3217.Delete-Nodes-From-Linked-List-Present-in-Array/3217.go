// https://leetcode.com/problems/delete-nodes-from-linked-list-present-in-array/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// ListNode is to define the struct
type ListNode = structures.ListNode

func modifiedList(nums []int, head *ListNode) *ListNode {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	dummy := &ListNode{}
	cur := dummy
	for head != nil {
		if _, exists := numSet[head.Val]; !exists {
			cur.Next = head
			cur = cur.Next
		}
		head = head.Next
	}
	cur.Next = nil
	return dummy.Next
}
