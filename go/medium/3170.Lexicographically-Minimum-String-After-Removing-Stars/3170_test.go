package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question3170 struct {
	params
	ans string
}

func Test_Problem3170(t *testing.T) {
	qs := []question3170{
		{
			params{s: "aaba*"},
			"aab",
		},
		{
			params{s: "abc"},
			"abc",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, clearStars(p.s), ans)
	}
}
