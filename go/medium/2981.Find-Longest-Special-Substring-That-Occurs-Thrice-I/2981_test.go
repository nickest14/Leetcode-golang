package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question2981 struct {
	params
	ans int
}

func Test_Problem2981(t *testing.T) {
	qs := []question2981{
		{
			params{s: "abcaba"},
			1,
		},
		{
			params{s: "aaaa"},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumLength(p.s), ans)
	}
}
