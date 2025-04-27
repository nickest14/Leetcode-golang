package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	minK int
	maxK int
}

type question2444 struct {
	params
	ans int64
}

func Test_Problem2444(t *testing.T) {
	qs := []question2444{
		{
			params{nums: []int{1, 3, 5, 2, 7, 5}, minK: 1, maxK: 5},
			2,
		},
		{
			params{nums: []int{1, 1, 1, 1}, minK: 1, maxK: 1},
			10,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countSubarrays(p.nums, p.minK, p.maxK), ans)
	}
}
