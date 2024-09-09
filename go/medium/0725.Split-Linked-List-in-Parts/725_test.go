package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	head *structures.ListNode
	k    int
}

type question725 struct {
	params
	ans []*structures.ListNode
}

func Test_Problem725(t *testing.T) {
	qs := []question725{
		{
			params{structures.IntsToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 3},
			[]*structures.ListNode{
				structures.IntsToList([]int{1, 2, 3, 4}),
				structures.IntsToList([]int{5, 6, 7}),
				structures.IntsToList([]int{8, 9, 10}),
			},
		},
		{
			params{structures.IntsToList([]int{1, 2, 3}), 5},
			[]*structures.ListNode{
				structures.IntsToList([]int{1}),
				structures.IntsToList([]int{2}),
				structures.IntsToList([]int{3}),
				structures.IntsToList([]int{}),
				structures.IntsToList([]int{}),
			},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		output := splitListToParts(p.head, p.k)
		for i := 0; i < len(output); i++ {
			fmt.Printf("index: %d        [optput]: %v    [answer]:%v\n", i, structures.ListToInts(output[i]), structures.ListToInts(ans[i]))
		}
	}
}
