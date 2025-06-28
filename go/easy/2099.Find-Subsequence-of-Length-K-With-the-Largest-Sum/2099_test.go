package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question2099 struct {
	params
	ans []int
}

func Test_Problem2099(t *testing.T) {
	qs := []question2099{
		{
			params{nums: []int{2, 1, 3, 3}, k: 2},
			[]int{3, 3},
		},
		{
			params{nums: []int{-1, -2, 3, 4}, k: 3},
			[]int{-1, 3, 4},
		},
		{
			params{nums: []int{3, 4, 3, 3}, k: 2},
			[]int{4, 3},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxSubsequence(p.nums, p.k), ans)
	}
}
