package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums    []int
	queries [][]int
}

type question3152 struct {
	params
	ans []bool
}

func Test_Problem3152(t *testing.T) {
	qs := []question3152{
		{
			params{nums: []int{3, 4, 1, 2, 6}, queries: [][]int{{0, 4}}},
			[]bool{false},
		},
		{
			params{nums: []int{4, 3, 1, 6}, queries: [][]int{{0, 2}, {2, 3}}},
			[]bool{false, true},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isArraySpecial(p.nums, p.queries), ans)
	}
}
