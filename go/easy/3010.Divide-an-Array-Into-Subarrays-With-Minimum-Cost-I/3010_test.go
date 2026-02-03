package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3010 struct {
	params
	ans int
}

func Test_Problem3010(t *testing.T) {
	qs := []question3010{
		{
			params{nums: []int{1, 2, 3, 12}},
			6,
		},
		{
			params{nums: []int{5, 4, 3}},
			12,
		},
		{
			params{nums: []int{10, 3, 1, 1}},
			12,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumCost(p.nums), ans)
	}
}
