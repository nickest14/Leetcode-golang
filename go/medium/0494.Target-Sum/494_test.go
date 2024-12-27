package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums   []int
	target int
}

type question494 struct {
	params
	ans int
}

func Test_Problem494(t *testing.T) {

	qs := []question494{
		{
			params{nums: []int{1, 1, 1, 1, 1}, target: 3},
			5,
		},
		{
			params{nums: []int{1}, target: 1},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, findTargetSumWays(p.nums, p.target), ans)
	}
}
