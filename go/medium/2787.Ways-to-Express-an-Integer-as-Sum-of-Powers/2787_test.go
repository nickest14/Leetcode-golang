package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
	x int
}

type question2787 struct {
	params
	ans int64
}

func Test_Problem2787(t *testing.T) {
	qs := []question2787{
		{
			params{n: 10, x: 2},
			1,
		},
		{
			params{n: 4, x: 1},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numberOfWays(p.n, p.x), ans)
	}
}
