package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question594 struct {
	params
	ans int
}

func Test_Problem594(t *testing.T) {
	qs := []question594{
		{
			params{[]int{1, 3, 2, 2, 5, 2, 3, 7}},
			5,
		},
		{
			params{[]int{1, 2, 3, 4}},
			2,
		},
		{
			params{[]int{1, 1, 1, 1}},
			0,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findLHS(p.nums), ans)
	}
}
