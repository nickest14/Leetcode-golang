package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3640 struct {
	params
	ans int64
}

func Test_Problem3640(t *testing.T) {
	qs := []question3640{
		{
			params{nums: []int{0, -2, -1, -3, 0, 2, -1}},
			-4,
		},
		{
			params{nums: []int{1, 4, 2, 7}},
			14,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxSumTrionic(p.nums), ans)
	}
}
