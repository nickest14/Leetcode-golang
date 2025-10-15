package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question3349 struct {
	params
	ans bool
}

func Test_Problem3349(t *testing.T) {
	qs := []question3349{
		{
			params{nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, k: 3},
			true,
		},
		{
			params{nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, k: 5},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, hasIncreasingSubarrays(p.nums, p.k), ans)
	}
}
