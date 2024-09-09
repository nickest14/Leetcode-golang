// https://leetcode.com/problems/split-linked-list-in-parts/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// ListNode is to define the struct
type ListNode = structures.ListNode

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	var ans []*ListNode
	baseSize, extra := length/k, length%k
	current = head
	for i := 0; i < k; i++ {
		partSize := baseSize
		if extra > 0 {
			extra--
			partSize++
		}
		var partHead, partTail *ListNode
		for j := 0; j < partSize; j++ {
			if partHead == nil {
				partHead, partTail = current, current
			} else {
				partTail.Next = current
				partTail = partTail.Next
			}
			if current != nil {
				current = current.Next
			}
		}
		if partTail != nil {
			partTail.Next = nil
		}
		ans = append(ans, partHead)
	}
	return ans
}
