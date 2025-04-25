package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question1399 struct {
	params
	ans int
}

func Test_Problem1399(t *testing.T) {
	qs := []question1399{
		{
			params{n: 13},
			4,
		},
		{
			params{n: 2},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countLargestGroup(p.n), ans)
	}
}
