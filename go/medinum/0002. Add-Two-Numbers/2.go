// https://leetcode.com/problems/add-two-numbers/
// Level: Medium

package leetcode

import (
	template "leetcode/go/template"
)

// ListNode is to define the struct
type ListNode = template.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	v1, v2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}
		current.Next = &ListNode{Val: (v1 + v2 + carry) % 10}
		carry = (v1 + v2 + carry) / 10
		current = current.Next
	}
	return head.Next
}
