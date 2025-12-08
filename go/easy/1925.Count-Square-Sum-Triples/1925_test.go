package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question1925 struct {
	params
	ans int
}

func Test_Problem1925(t *testing.T) {
	qs := []question1925{
		{
			params{n: 5},
			2,
		},
		{
			params{n: 10},
			4,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countTriples(p.n), ans)
	}
}
