package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
	k int
}

type question1945 struct {
	params
	ans int
}

func Test_Problem1945(t *testing.T) {
	qs := []question1945{
		{
			params{"leetcode", 2},
			6,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, getLucky(p.s, p.k), ans)
	}
}
