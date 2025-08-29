package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
	m int
}

type question3021 struct {
	params
	ans int
}

func Test_Problem3021(t *testing.T) {
	qs := []question3021{
		{
			params{n: 3, m: 2},
			3,
		},
		{
			params{n: 1, m: 1},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, flowerGame(p.n, p.m), ans)
	}
}
