package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question2906 struct {
	params
	ans [][]int
}

func Test_Problem2906(t *testing.T) {
	qs := []question2906{
		{
			params{grid: [][]int{{1, 2}, {3, 4}}},
			[][]int{{24, 12}, {8, 6}},
		},
		{
			params{grid: [][]int{{12345}, {2}, {1}}},
			[][]int{{2}, {0}, {0}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, constructProductMatrix(p.grid), ans)
	}
}
