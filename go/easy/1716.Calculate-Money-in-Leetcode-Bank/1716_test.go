package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question1716 struct {
	params
	ans int
}

func Test_Problem1716(t *testing.T) {
	qs := []question1716{
		{
			params{4},
			10,
		},
		{
			params{10},
			37,
		},
		{
			params{20},
			96,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, totalMoney(p.n), ans)
	}
}
