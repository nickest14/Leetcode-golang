package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question790 struct {
	params
	ans int
}

func Test_Problem790(t *testing.T) {
	qs := []question790{
		{
			params{3},
			5,
		},
		{
			params{1},
			1,
		},
		{
			params{5},
			24,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numTilings(p.n), ans)
	}
}
