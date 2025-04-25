package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums   []int
	modulo int
	k      int
}

type question2707 struct {
	params
	ans int64
}

func Test_Problem2707(t *testing.T) {
	qs := []question2707{
		{
			params{nums: []int{3, 2, 4}, modulo: 2, k: 1},
			3,
		},
		{
			params{nums: []int{3, 1, 9, 6}, modulo: 3, k: 0},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countInterestingSubarrays(p.nums, p.modulo, p.k), ans)
	}
}
