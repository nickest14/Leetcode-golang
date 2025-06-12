package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question3442 struct {
	params
	ans int
}

func Test_Problem3442(t *testing.T) {
	qs := []question3442{
		{
			params{s: "aaaaabbc"},
			3,
		},
		{
			params{s: "abcabcab"},
			1,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxDifference(p.s), ans)
	}
}
