package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
	k int
}

type question440 struct {
	params
	ans int
}

func Test_Problem440(t *testing.T) {
	qs := []question440{
		{
			params{n: 13, k: 2},
			10,
		},
		{
			params{n: 1, k: 1},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findKthNumber(p.n, p.k), ans)
	}
}
