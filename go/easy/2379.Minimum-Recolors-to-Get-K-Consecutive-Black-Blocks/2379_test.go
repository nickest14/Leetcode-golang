package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	blocks string
	k      int
}

type question2379 struct {
	params
	ans int
}

func Test_Problem2379(t *testing.T) {
	qs := []question2379{
		{
			params{blocks: "WBBWWBBWBW", k: 7},
			3,
		},
		{
			params{blocks: "WBWBBBW", k: 2},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumRecolors(p.blocks, p.k), ans)
	}
}
