package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2221 struct {
	params
	ans int
}

func Test_Problem2221(t *testing.T) {
	qs := []question2221{
		{
			params{nums: []int{1, 2, 3, 4, 5}},
			8,
		},
		{
			params{nums: []int{5}},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, triangularSum(p.nums), ans)
	}
}
