package structures

// ListNode is to define linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListToInts convert List to []int
func ListToInts(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// IntsToList convert []int to List
func IntsToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
