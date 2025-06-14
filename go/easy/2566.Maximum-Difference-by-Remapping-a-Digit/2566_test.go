package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	num int
}

type question2843 struct {
	params
	ans int
}

func Test_Problem2843(t *testing.T) {
	qs := []question2843{
		{
			params{11891},
			99009,
		},
		{
			params{90},
			99,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minMaxDifference(p.num), ans)
	}
}
