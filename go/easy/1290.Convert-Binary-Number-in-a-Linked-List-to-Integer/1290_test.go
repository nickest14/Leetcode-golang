package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	head *structures.ListNode
}

type question1290 struct {
	params
	ans int
}

func Test_Problem1290(t *testing.T) {
	qs := []question1290{
		{
			params{structures.IntsToList([]int{1, 0, 1})},
			5,
		},
		{
			params{structures.IntsToList([]int{0})},
			0,
		},
	}
	for _, q := range qs {
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", q.params, getDecimalValue(q.params.head), q.ans)
	}
}
