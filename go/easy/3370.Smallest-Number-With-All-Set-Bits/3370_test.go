package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question3370 struct {
	params
	ans int
}

func Test_Problem3370(t *testing.T) {
	qs := []question3370{
		{
			params{5},
			7,
		},
		{
			params{10},
			15,
		},
		{
			params{3},
			3,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, smallestNumber(p.n), ans)
	}
}
