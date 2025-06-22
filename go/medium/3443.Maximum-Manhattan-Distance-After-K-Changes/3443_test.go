package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
	k int
}

type question3443 struct {
	params
	ans int
}

func Test_Problem3443(t *testing.T) {
	qs := []question3443{
		{
			params{s: "NWSE", k: 1},
			3,
		},
		{
			params{s: "NSWWEW", k: 3},
			6,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxDistance(p.s, p.k), ans)
	}
}
