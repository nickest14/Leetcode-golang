package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	points [][]int
}

type question812 struct {
	params
	ans float64
}

func Test_Problem812(t *testing.T) {

	qs := []question812{
		{
			params{[][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}},
			2.00000,
		},
		{
			params{[][]int{{1, 0}, {0, 0}, {0, 1}}},
			0.50000,
		},
	}

	for _, q := range qs {
		p := q.params
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, largestTriangleArea(p.points), q.ans)
	}
}
