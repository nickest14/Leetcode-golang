package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question1653 struct {
	params
	ans int
}

func Test_Problem1653(t *testing.T) {
	qs := []question1653{
		{
			params{"aababbab"},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumDeletions(p.s), ans)
	}
}
