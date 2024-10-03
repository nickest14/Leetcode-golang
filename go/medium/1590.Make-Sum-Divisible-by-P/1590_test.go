package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	p    int
}

type question1590 struct {
	params
	ans int
}

func Test_Problem(t *testing.T) {
	qs := []question1590{
		// {
		// 	params{nums: []int{3, 1, 4, 2}, p: 6},
		// 	1,
		// },
		// {
		// 	params{nums: []int{6, 3, 5, 2}, p: 9},
		// 	2,
		// },
		{
			params{nums: []int{2, 3, 11}, p: 5},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minSubarray(p.nums, p.p), ans)
	}
}
