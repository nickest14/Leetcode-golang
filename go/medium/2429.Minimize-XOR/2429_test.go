package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	num1 int
	num2 int
}

type question2429 struct {
	params
	ans int
}

func Test_Problem2429(t *testing.T) {
	qs := []question2429{
		{
			params{num1: 3, num2: 5},
			3,
		},
		{
			params{num1: 1, num2: 12},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimizeXor(p.num1, p.num2), ans)
	}
}
