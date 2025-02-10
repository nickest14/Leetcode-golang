package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question3174 struct {
	params
	ans string
}

func Test_Problem3174(t *testing.T) {
	qs := []question3174{
		{
			params{s: "abc"},
			"abc",
		},
		{
			params{s: "cb34"},
			"",
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, clearDigits(p.s), ans)
	}
}
