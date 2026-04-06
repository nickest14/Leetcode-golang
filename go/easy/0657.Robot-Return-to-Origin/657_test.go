package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	moves string
}

type question657 struct {
	params
	ans bool
}

func Test_Problem657(t *testing.T) {
	qs := []question657{
		{
			params{moves: "UD"},
			true,
		},
		{
			params{moves: "LL"},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, judgeCircle(p.moves), ans)
	}
}
