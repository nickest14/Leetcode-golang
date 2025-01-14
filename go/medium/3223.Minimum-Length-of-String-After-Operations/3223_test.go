package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question3223 struct {
	params
	ans int
}

func Test_Problem3223(t *testing.T) {
	qs := []question3223{
		{
			params{s: "abaacbcbb"},
			5,
		},
		{
			params{s: "aa"},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumLength(p.s), ans)
	}
}
