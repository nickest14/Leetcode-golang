package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	m      int
	n      int
	guards [][]int
	walls  [][]int
}

type question2257 struct {
	params
	ans int
}

func Test_Problem2257(t *testing.T) {
	qs := []question2257{
		{
			params{m: 4, n: 6, guards: [][]int{{0, 0}, {1, 1}, {2, 3}}, walls: [][]int{{0, 1}, {2, 2}, {1, 4}}},
			7,
		},
		{
			params{m: 3, n: 3, guards: [][]int{{1, 1}}, walls: [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countUnguarded(p.m, p.n, p.guards, p.walls), ans)
	}
}
