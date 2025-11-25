package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	k int
}

type question1015 struct {
	params
	ans int
}

func Test_Problem1015(t *testing.T) {
	qs := []question1015{
		{
			params{1},
			1,
		},
		{
			params{2},
			-1,
		},
		{
			params{3},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, smallestRepunitDivByK(p.k), ans)
	}
}
