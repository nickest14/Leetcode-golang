package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums  []int
	lower int
	upper int
}

type question2563 struct {
	params
	ans int
}

func Test_Problem2563(t *testing.T) {
	qs := []question2563{
		{
			params{nums: []int{0, 1, 7, 4, 4, 5}, lower: 3, upper: 6},
			6,
		},
		{
			params{nums: []int{1, 7, 9, 2, 5}, lower: 11, upper: 11},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countFairPairs(p.nums, p.lower, p.upper), ans)
	}
}
