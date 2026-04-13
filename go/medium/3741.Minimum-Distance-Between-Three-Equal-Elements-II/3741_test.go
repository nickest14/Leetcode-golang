package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3741 struct {
	params
	ans int
}

func Test_Problem3741(t *testing.T) {
	qs := []question3741{
		{
			params{nums: []int{1, 2, 1, 1, 3}},
			6,
		},
		{
			params{nums: []int{1, 1, 2, 3, 2, 1, 2}},
			8,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumDistance(p.nums), ans)
	}
}
