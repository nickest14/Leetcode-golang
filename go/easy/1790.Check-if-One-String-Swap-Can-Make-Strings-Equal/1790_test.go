package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s1 string
	s2 string
}

type question1790 struct {
	params
	ans bool
}

func Test_Problem1790(t *testing.T) {
	qs := []question1790{
		{
			params{"bank", "kanb"},
			true,
		},
		{
			params{"attack", "defend"},
			false,
		},
		{
			params{"kelb", "kelb"},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, areAlmostEqual(p.s1, p.s2), ans)
	}
}
