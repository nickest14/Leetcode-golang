// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
// Level: Easy

package leetcode

import (
	structures "leetcode/go/structures"
	"strconv"
)

// ListNode is to define the struct
type ListNode = structures.ListNode

func getDecimalValue(head *ListNode) int {
	s := ""
	for head != nil {
		s += strconv.Itoa(head.Val)
		head = head.Next
	}
	ans, _ := strconv.ParseInt(s, 2, 64)
	return int(ans)
}
