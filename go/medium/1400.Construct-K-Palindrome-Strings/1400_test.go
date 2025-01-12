package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
	k int
}

type question1400 struct {
	params
	ans bool
}

func Test_Problem1400(t *testing.T) {
	qs := []question1400{
		{
			params{s: "annabelle", k: 2},
			true,
		},
		{
			params{s: "leetcode", k: 3},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, canConstruct(p.s, p.k), ans)
	}
}
