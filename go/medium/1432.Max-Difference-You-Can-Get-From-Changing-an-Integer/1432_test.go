package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	num int
}

type question1432 struct {
	params
	ans int
}

func Test_Problem1432(t *testing.T) {
	qs := []question1432{
		{
			params{555},
			888,
		},
		{
			params{9},
			8,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxDiff(p.num), ans)
	}
}
