package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	points [][]int
}

type question3027 struct {
	params
	ans int
}

func Test_Problem3027(t *testing.T) {
	qs := []question3027{
		{
			params{points: [][]int{{1, 1}, {2, 2}, {3, 3}}},
			0,
		},
		{
			params{points: [][]int{{6, 2}, {4, 4}, {2, 6}}},
			2,
		},
		{
			params{points: [][]int{{3, 1}, {1, 3}, {1, 1}}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numberOfPairs(p.points), ans)
	}
}
