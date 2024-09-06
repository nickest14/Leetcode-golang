package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	nums []int
	head *structures.ListNode
}

type question3217 struct {
	params
	ans *structures.ListNode
}

func Test_Problem3217(t *testing.T) {
	qs := []question3217{
		{
			params{[]int{1, 2, 3}, structures.IntsToList([]int{1, 2, 3, 4})},
			structures.IntsToList([]int{4}),
		},
		{
			params{[]int{1}, structures.IntsToList([]int{1, 2, 1, 2, 1, 2, 1, 2})},
			structures.IntsToList([]int{2, 2, 2, 2}),
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, modifiedList(p.nums, p.head), ans)
	}
}
