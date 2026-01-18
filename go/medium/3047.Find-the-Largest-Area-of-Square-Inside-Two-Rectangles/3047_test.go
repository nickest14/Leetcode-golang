package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	bottomLeft [][]int
	topRight   [][]int
}

type question3025 struct {
	params
	ans int
}

func Test_Problem3025(t *testing.T) {
	qs := []question3025{
		{
			params{bottomLeft: [][]int{{1, 1}, {2, 2}, {3, 1}}, topRight: [][]int{{3, 3}, {4, 4}, {6, 6}}},
			1,
		},
		{
			params{bottomLeft: [][]int{{1, 1}, {1, 3}, {1, 5}}, topRight: [][]int{{5, 5}, {5, 7}, {5, 9}}},
			4,
		},
		{
			params{bottomLeft: [][]int{{1, 1}, {2, 2}, {1, 2}}, topRight: [][]int{{3, 3}, {4, 4}, {3, 4}}},
			1,
		},
		{
			params{bottomLeft: [][]int{{1, 1}, {3, 3}, {3, 1}}, topRight: [][]int{{2, 2}, {4, 4}, {4, 2}}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, largestSquareArea(p.bottomLeft, p.topRight), ans)
	}
}
