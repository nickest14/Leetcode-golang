package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	points [][]int
}

type question3623 struct {
	params
	ans int
}

func Test_Problem3623(t *testing.T) {
	qs := []question3623{
		{
			params{[][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}}},
			3,
		},
		{
			params{[][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countTrapezoids(p.points), ans)
	}
}
