package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question3227 struct {
	params
	ans bool
}

func Test_Problem3227(t *testing.T) {
	qs := []question3227{
		{
			params{"leetcoder"},
			true,
		},
		{
			params{"bbcd"},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, doesAliceWin(p.s), ans)
	}
}
