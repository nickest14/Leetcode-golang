package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question3381 struct {
	params
	ans int64
}

func Test_Problem3381(t *testing.T) {
	qs := []question3381{
		{
			params{nums: []int{1, 2}, k: 1},
			3,
		},
		{
			params{nums: []int{-1, -2, -3, -4, -5}, k: 4},
			-10,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxSubarraySum(p.nums, p.k), ans)
	}
}
