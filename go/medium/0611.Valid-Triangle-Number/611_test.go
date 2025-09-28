package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question611 struct {
	params
	ans int
}

func Test_Problem611(t *testing.T) {
	qs := []question611{
		{
			params{nums: []int{2, 2, 3, 4}},
			3,
		},
		{
			params{nums: []int{4, 2, 3, 4}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, triangleNumber(p.nums), ans)
	}
}
