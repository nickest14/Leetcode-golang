package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3379 struct {
	params
	ans []int
}

func Test_Problem3379(t *testing.T) {
	qs := []question3379{
		{
			params{nums: []int{3, -2, 1, 1}},
			[]int{1, 1, 1, 3},
		},
		{
			params{nums: []int{-1, 4, -1}},
			[]int{-1, -1, 4},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, constructTransformedArray(p.nums), ans)
	}
}
