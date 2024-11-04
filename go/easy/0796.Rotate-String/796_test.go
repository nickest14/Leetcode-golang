package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s    string
	goal string
}

type question796 struct {
	params
	ans bool
}

func Test_Problem796(t *testing.T) {
	qs := []question796{
		{
			params{s: "abcde", goal: "cdeab"},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, rotateString(p.s, p.goal), ans)
	}
}
