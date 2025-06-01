package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n     int
	limit int
}

type question2929 struct {
	params
	ans int64
}

func Test_Problem2929(t *testing.T) {
	qs := []question2929{
		{
			params{n: 5, limit: 2},
			3,
		},
		{
			params{n: 3, limit: 3},
			10,
		},
		{
			params{n: 4, limit: 1},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, distributeCandies(p.n, p.limit), ans)
	}
}
