package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3487 struct {
	params
	ans int
}

func Test_Problem3487(t *testing.T) {
	qs := []question3487{
		{
			params{nums: []int{1, 2, 3, 4, 5}},
			15,
		},
		{
			params{nums: []int{1, 1, 0, 1, 1}},
			1,
		},
		{
			params{nums: []int{1, 2, -1, -2, 1, 0, -1}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxSum(p.nums), ans)
	}
}
