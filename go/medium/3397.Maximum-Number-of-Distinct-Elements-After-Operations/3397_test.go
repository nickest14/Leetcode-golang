package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question3397 struct {
	params
	ans int
}

func Test_Problem3397(t *testing.T) {
	qs := []question3397{
		{
			params{nums: []int{1, 2, 2, 3, 3, 4}, k: 2},
			6,
		},
		{
			params{nums: []int{4, 4, 4, 4}, k: 1},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxDistinctElements(p.nums, p.k), ans)
	}
}
