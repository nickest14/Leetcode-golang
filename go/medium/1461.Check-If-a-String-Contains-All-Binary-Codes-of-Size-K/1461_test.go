package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
	k int
}

type question1461 struct {
	params
	ans bool
}

func Test_Problem1461(t *testing.T) {
	qs := []question1461{
		{
			params{s: "00110110", k: 2},
			true,
		},
		{
			params{s: "0110", k: 1},
			true,
		},
		{
			params{s: "0110", k: 2},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, hasAllCodes(p.s, p.k), ans)
	}
}
