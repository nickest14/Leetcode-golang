package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question3541 struct {
	params
	ans int
}

func Test_Problem3541(t *testing.T) {
	qs := []question3541{
		{
			params{25},
			27,
		},
		{
			params{10},
			9,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, mirrorDistance(p.n), ans)
	}
}
