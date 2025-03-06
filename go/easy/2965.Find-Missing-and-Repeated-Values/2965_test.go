package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question2965 struct {
	params
	ans []int
}

func Test_Problem2965(t *testing.T) {
	qs := []question2965{
		{
			params{[][]int{{1, 3}, {2, 2}}},
			[]int{2, 4},
		},
		{
			params{[][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}},
			[]int{9, 5},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findMissingAndRepeatedValues(p.grid), ans)
	}
}
