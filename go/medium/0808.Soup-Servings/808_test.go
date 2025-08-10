package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question808 struct {
	params
	ans float64
}

func Test_Problem808(t *testing.T) {
	qs := []question808{
		{
			params{50},
			0.62500,
		},
		{
			params{100},
			0.71875,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, soupServings(p.n), ans)
	}
}
