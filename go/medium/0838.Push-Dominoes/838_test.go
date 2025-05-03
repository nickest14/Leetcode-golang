package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	dominoes string
}

type question838 struct {
	params
	ans string
}

func Test_Problem838(t *testing.T) {
	qs := []question838{
		{
			params{"RR.L"},
			"RR.L",
		},
		{
			params{".L.R...LR..L.."},
			"LL.RR.LLRRLL..",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, pushDominoes(p.dominoes), ans)
	}
}
