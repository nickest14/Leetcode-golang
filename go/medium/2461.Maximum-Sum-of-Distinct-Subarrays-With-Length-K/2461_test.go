package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question2461 struct {
	params
	ans int
}

func Test_Problem2461(t *testing.T) {
	qs := []question2461{
		{
			params{nums: []int{1, 5, 4, 2, 9, 9, 9}, k: 3},
			15,
		},
		{
			params{nums: []int{4, 4, 4}, k: 3},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumSubarraySum(p.nums, p.k), ans)
	}
}
