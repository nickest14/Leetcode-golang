package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1749 struct {
	params
	ans int
}

func Test_Problem1749(t *testing.T) {

	qs := []question1749{
		{
			params{[]int{1, -3, 2, 3, -4}},
			5,
		},
		{
			params{[]int{2, -5, 1, -4, 3, -2}},
			8,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, maxAbsoluteSum(p.nums), ans)
	}
}
