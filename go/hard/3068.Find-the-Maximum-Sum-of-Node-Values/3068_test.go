package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums  []int
	k     int
	edges [][]int
}

type question3068 struct {
	params
	ans int
}

func Test_Problem3068(t *testing.T) {
	qs := []question3068{
		{
			params{nums: []int{1, 2, 1}, k: 3, edges: [][]int{{0, 1}, {0, 2}}},
			6,
		},
		{
			params{nums: []int{2, 3}, k: 7, edges: [][]int{{0, 1}}},
			9,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumValueSum(p.nums, p.k, p.edges), ans)
	}
}
