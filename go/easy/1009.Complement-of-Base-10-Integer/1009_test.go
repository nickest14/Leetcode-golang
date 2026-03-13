package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question1009 struct {
	params
	ans int
}

func Test_Problem1009(t *testing.T) {
	qs := []question1009{
		{
			params{5},
			2,
		},
		{
			params{7},
			0,
		},
		{
			params{10},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, bitwiseComplement(p.n), ans)
	}
}
