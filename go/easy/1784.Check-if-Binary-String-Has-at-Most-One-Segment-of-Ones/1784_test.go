package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1784 struct {
	params
	ans bool
}

func Test_Problem1784(t *testing.T) {
	qs := []question1784{
		{
			params{s: "1001"},
			false,
		},
		{
			params{s: "110"},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, checkOnesSegment(p.s), ans)
	}
}
