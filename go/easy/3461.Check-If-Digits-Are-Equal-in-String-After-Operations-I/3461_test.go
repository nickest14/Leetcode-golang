package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question3461 struct {
	params
	ans bool
}

func Test_Problem3461(t *testing.T) {
	qs := []question3461{
		{
			params{s: "3902"},
			true,
		},
		{
			params{s: "34789"},
			false,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, hasSameDigits(p.s), ans)
	}
}
