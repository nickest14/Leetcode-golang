package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	num1 int
	num2 int
}

type question2749 struct {
	params
	ans int
}

func Test_Problem2749(t *testing.T) {
	qs := []question2749{
		{
			params{num1: 3, num2: -2},
			3,
		},
		{
			params{num1: 5, num2: 7},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, makeTheIntegerZero(p.num1, p.num2), ans)
	}
}
