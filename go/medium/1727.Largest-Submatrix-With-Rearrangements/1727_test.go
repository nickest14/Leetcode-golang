package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	matrix [][]int
}

type question2110 struct {
	params
	ans int
}

func Test_Problem2110(t *testing.T) {
	qs := []question2110{
		{
			params{matrix: [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}},
			4,
		},
		{
			params{matrix: [][]int{{1, 0, 1, 0, 1}}},
			3,
		},
		{
			params{matrix: [][]int{{1, 1, 0}, {1, 0, 1}}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, largestSubmatrix(p.matrix), ans)
	}
}
