package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	low  int
	high int
}

type question2843 struct {
	params
	ans int
}

func Test_Problem2843(t *testing.T) {
	qs := []question2843{
		{
			params{low: 1, high: 100},
			9,
		},
		{
			params{low: 1200, high: 1230},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countSymmetricIntegers(p.low, p.high), ans)
	}
}
