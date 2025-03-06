package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question2579 struct {
	params
	ans int
}

func Test_Problem2579(t *testing.T) {
	qs := []question2579{
		{
			params{2},
			5,
		},
		{
			params{3},
			13,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, coloredCells(p.n), ans)
	}
}
