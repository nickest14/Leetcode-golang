package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1358 struct {
	params
	ans int
}

func Test_Problem1358(t *testing.T) {
	qs := []question1358{
		{
			params{"abcabc"},
			10,
		},
		{
			params{"aaacb"},
			3,
		},
		{
			params{"abc"},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numberOfSubstrings(p.s), ans)
	}
}
