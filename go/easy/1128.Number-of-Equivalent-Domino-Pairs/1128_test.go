package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	dominoes [][]int
}

type question1128 struct {
	params
	ans int
}

func Test_Problem1128(t *testing.T) {

	qs := []question1128{
		{
			params{[][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}},
			1,
		},
		{
			params{[][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}},
			3,
		},
	}

	for _, q := range qs {
		p := q.params
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numEquivDominoPairs(p.dominoes), q.ans)
	}
}
