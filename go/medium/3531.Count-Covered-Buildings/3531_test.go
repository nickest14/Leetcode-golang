package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n         int
	buildings [][]int
}

type question3531 struct {
	params
	ans int
}

func Test_Problem3531(t *testing.T) {
	qs := []question3531{
		{
			params{n: 3, buildings: [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}}},
			1,
		},
		{
			params{n: 3, buildings: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
			0,
		},
		{
			params{n: 5, buildings: [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countCoveredBuildings(p.n, p.buildings), ans)
	}
}
