package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int64
}

type question1922 struct {
	params
	ans int
}

func Test_Problem1922(t *testing.T) {
	qs := []question1922{
		{
			params{1},
			5,
		},
		{
			params{4},
			400,
		},
		{
			params{50},
			564908303,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countGoodNumbers(p.n), ans)
	}
}
