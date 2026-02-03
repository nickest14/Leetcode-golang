package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3637 struct {
	params
	ans bool
}

func Test_Problem3637(t *testing.T) {
	qs := []question3637{
		{
			params{nums: []int{1, 3, 5, 4, 2, 6}},
			true,
		},
		{
			params{nums: []int{2, 1, 3}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isTrionic(p.nums), ans)
	}
}
