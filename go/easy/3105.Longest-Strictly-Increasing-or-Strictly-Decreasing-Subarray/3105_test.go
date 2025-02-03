package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3105 struct {
	params
	ans int
}

func Test_Problem3105(t *testing.T) {
	qs := []question3105{
		{
			params{[]int{1, 4, 3, 3, 2}},
			2,
		},
		{
			params{[]int{3, 3, 3, 3}},
			1,
		},
		{
			params{[]int{1, 2, 3}},
			3,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, longestMonotonicSubarray(p.nums), ans)
	}
}
