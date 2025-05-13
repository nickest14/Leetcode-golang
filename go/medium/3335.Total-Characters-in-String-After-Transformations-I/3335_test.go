package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
	t int
}

type question3335 struct {
	params
	ans int
}

func Test_Problem3335(t *testing.T) {
	qs := []question3335{
		{
			params{s: "abcyy", t: 2},
			7,
		},
		{
			params{s: "azbk", t: 1},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, lengthAfterTransformations(p.s, p.t), ans)
	}
}
