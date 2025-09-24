package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	numerator   int
	denominator int
}

type question166 struct {
	params
	ans string
}

func Test_Problem166(t *testing.T) {
	qs := []question166{
		{
			params{numerator: 1, denominator: 2},
			"0.5",
		},
		{
			params{numerator: 2, denominator: 1},
			"2",
		},
		{
			params{numerator: 4, denominator: 333},
			"0.(012)",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, fractionToDecimal(p.numerator, p.denominator), ans)
	}
}
