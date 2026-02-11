package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3719 struct {
	params
	ans int
}

func Test_Problem3719(t *testing.T) {
	qs := []question3719{
		{
			params{nums: []int{2, 5, 4, 3}},
			4,
		},
		{
			params{nums: []int{3, 2, 2, 5, 4}},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, longestBalanced(p.nums), ans)
	}
}
