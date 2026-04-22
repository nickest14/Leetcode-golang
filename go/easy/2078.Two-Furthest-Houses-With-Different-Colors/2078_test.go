package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	colors []int
}

type question2078 struct {
	params
	ans int
}

func Test_Problem2078(t *testing.T) {
	qs := []question2078{
		{
			params{[]int{1, 1, 1, 6, 1, 1, 1}},
			3,
		},
		{
			params{[]int{1, 8, 3, 8, 3}},
			4,
		},
		{
			params{[]int{0, 1}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxDistance(p.colors), ans)
	}
}
