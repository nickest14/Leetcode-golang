package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
	m int
}

type question2894 struct {
	params
	ans int
}

func Test_Problem2894(t *testing.T) {
	qs := []question2894{
		{
			params{n: 10, m: 3},
			19,
		},
		{
			params{n: 5, m: 6},
			15,
		},
		{
			params{n: 5, m: 1},
			-15,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, differenceOfSums(p.n, p.m), ans)
	}
}
