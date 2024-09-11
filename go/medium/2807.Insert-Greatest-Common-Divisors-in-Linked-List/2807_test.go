package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	head *structures.ListNode
}

type question2326 struct {
	params
	ans *structures.ListNode
}

func Test_Problem2326(t *testing.T) {
	qs := []question2326{
		{
			params{structures.IntsToList([]int{18, 6, 10, 3})},
			structures.IntsToList([]int{18, 6, 6, 2, 10, 1, 3}),
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, structures.ListToInts(insertGreatestCommonDivisors(p.head)), structures.ListToInts(ans))
	}
}
