package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3423 struct {
	params
	ans int
}

func Test_Problem3423(t *testing.T) {
	qs := []question3423{
		{
			params{nums: []int{1, 2, 4}},
			3,
		},
		{
			params{nums: []int{-5, -10, -5}},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxAdjacentDistance(p.nums), ans)
	}
}
