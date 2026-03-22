package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	mat    [][]int
	target [][]int
}

type question1886 struct {
	params
	ans bool
}

func Test_Problem1886(t *testing.T) {
	qs := []question1886{
		{
			params{mat: [][]int{{0, 1}, {1, 0}}, target: [][]int{{1, 0}, {0, 1}}},
			true,
		},
		{
			params{mat: [][]int{{0, 1}, {1, 1}}, target: [][]int{{1, 0}, {0, 1}}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findRotation(p.mat, p.target), ans)
	}
}
