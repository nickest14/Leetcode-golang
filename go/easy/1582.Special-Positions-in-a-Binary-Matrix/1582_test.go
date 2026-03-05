package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	mat [][]int
}

type question1582 struct {
	params
	ans int
}

func Test_Problem1582(t *testing.T) {
	qs := []question1582{
		{
			params{mat: [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}},
			1,
		},
		{
			params{mat: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numSpecial(p.mat), ans)
	}
}
