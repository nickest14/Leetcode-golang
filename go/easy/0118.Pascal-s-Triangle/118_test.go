package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	numRows int
}

type question118 struct {
	params
	ans [][]int
}

func Test_Problem796(t *testing.T) {
	qs := []question118{
		{
			params{5},
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			params{1},
			[][]int{{1}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, generate(p.numRows), ans)
	}
}
