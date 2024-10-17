package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	num int
}

type question670 struct {
	params
	ans int
}

func Test_Problem670(t *testing.T) {
	qs := []question670{
		{
			params{2736},
			7236,
		},
		{
			params{9973},
			9973,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumSwap(p.num), ans)
	}
}
