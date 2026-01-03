package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question1411 struct {
	params
	ans int
}

func Test_Problem1411(t *testing.T) {
	qs := []question1411{
		{
			params{n: 1},
			12,
		},
		{
			params{5000},
			30228214,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numOfWays(p.n), ans)
	}
}
