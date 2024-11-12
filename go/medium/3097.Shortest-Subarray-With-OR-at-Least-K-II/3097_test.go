package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question3097 struct {
	params
	ans int
}

func Test_Problem3097(t *testing.T) {
	qs := []question3097{
		{
			params{nums: []int{1, 2, 3}, k: 2},
			1,
		},
		{
			params{nums: []int{2, 1, 8}, k: 10},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumSubarrayLength(p.nums, p.k), ans)
	}
}
