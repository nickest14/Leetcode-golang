package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	points [][]int
}

type question3025 struct {
	params
	ans int
}

func Test_Problem3025(t *testing.T) {
	qs := []question3025{
		{
			params{points: [][]int{{1, 1}, {2, 2}, {3, 3}}},
			0,
		},
		{
			params{points: [][]int{{6, 6}, {4, 4}, {2, 6}}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numberOfPairs(p.points), ans)
	}
}
