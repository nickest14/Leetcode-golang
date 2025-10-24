package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question2048 struct {
	params
	ans int
}

func Test_Problem2048(t *testing.T) {
	qs := []question2048{
		{
			params{n: 1},
			22,
		},
		{
			params{n: 1000},
			1333,
		},
		{
			params{n: 3000},
			3133,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, nextBeautifulNumber(p.n), ans)
	}
}
