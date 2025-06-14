package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	p    int
}

type question2616 struct {
	params
	ans int
}

func Test_Problem2616(t *testing.T) {
	qs := []question2616{
		{
			params{nums: []int{10, 1, 2, 7, 1, 3}, p: 2},
			1,
		},
		{
			params{nums: []int{4, 2, 1, 2}, p: 1},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimizeMax(p.nums, p.p), ans)
	}
}
