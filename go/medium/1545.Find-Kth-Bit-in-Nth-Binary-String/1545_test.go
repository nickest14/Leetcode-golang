package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
	k int
}

type question2044 struct {
	params
	ans byte
}

func Test_Problem2044(t *testing.T) {
	qs := []question2044{
		{
			params{n: 3, k: 1},
			'0',
		},
		{
			params{n: 4, k: 11},
			'1',
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findKthBit(p.n, p.k), ans)
	}
}
