package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2529 struct {
	params
	ans int
}

func Test_Problem2529(t *testing.T) {
	qs := []question2529{
		{
			params{[]int{-2, -1, -1, 1, 2, 3}},
			3,
		},
		{
			params{[]int{-3, -2, -1, 0, 0, 1, 2}},
			3,
		},
		{
			params{[]int{5, 20, 66, 1314}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumCount(p.nums), ans)
	}
}
