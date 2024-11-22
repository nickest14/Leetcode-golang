package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	matrix [][]int
}

type question1072 struct {
	params
	ans int
}

func Test_Problem1072(t *testing.T) {
	qs := []question1072{
		// {
		// 	params{[][]int{{0, 1}, {1, 0}}},
		// 	1,
		// },
		{
			params{[][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxEqualRowsAfterFlips(p.matrix), ans)
	}
}
