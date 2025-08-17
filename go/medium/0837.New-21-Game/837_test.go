package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n      int
	k      int
	maxPts int
}

type question837 struct {
	params
	ans float64
}

func Test_Problem837(t *testing.T) {
	qs := []question837{
		{
			params{n: 10, k: 1, maxPts: 10},
			1.00000,
		},
		{
			params{n: 6, k: 1, maxPts: 10},
			0.60000,
		},
		{
			params{n: 21, k: 17, maxPts: 10},
			0.73278,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, new21Game(p.n, p.k, p.maxPts), ans)
	}
}
