package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	mat       [][]int
	threshold int
}

type question1292 struct {
	params
	ans int
}

func Test_Problem1292(t *testing.T) {
	qs := []question1292{
		{
			params{mat: [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}, threshold: 4},
			2,
		},
		{
			params{mat: [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}, threshold: 1},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxSideLength(p.mat, p.threshold), ans)
	}
}
