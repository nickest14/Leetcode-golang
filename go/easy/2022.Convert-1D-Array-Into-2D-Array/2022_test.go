package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	original []int
	m        int
	n        int
}

type question2022 struct {
	params
	ans [][]int
}

func Test_Problem2022(t *testing.T) {
	qs := []question2022{
		{
			params{original: []int{1, 2, 3, 4}, m: 2, n: 2},
			[][]int{{1, 2}, {3, 4}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, construct2DArray(p.original, p.m, p.n), ans)
	}
}
