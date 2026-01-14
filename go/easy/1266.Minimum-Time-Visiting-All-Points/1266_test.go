package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	points [][]int
}

type question1266 struct {
	params
	ans int
}

func Test_Problem1266(t *testing.T) {

	qs := []question1266{
		{
			params{[][]int{{1, 1}, {3, 4}, {-1, 0}}},
			7,
		},
		{
			params{[][]int{{3, 2}, {-2, 2}}},
			5,
		},
	}

	for _, q := range qs {
		p := q.params
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minTimeToVisitAllPoints(p.points), q.ans)
	}
}
