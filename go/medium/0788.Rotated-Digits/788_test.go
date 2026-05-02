package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question788 struct {
	params
	ans int
}

func Test_Problem788(t *testing.T) {
	qs := []question788{
		{
			params{10},
			4,
		},
		{
			params{1},
			0,
		},
		{
			params{2},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, rotatedDigits(p.n), ans)
	}
}
