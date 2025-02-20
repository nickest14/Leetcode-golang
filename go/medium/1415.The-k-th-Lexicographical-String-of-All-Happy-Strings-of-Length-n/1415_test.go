package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
	k int
}

type question1415 struct {
	params
	ans string
}

func Test_Problem1415(t *testing.T) {
	qs := []question1415{
		{
			params{n: 1, k: 3},
			"c",
		},
		{
			params{n: 3, k: 9},
			"cab",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, getHappyString(p.n, p.k), ans)
	}
}
