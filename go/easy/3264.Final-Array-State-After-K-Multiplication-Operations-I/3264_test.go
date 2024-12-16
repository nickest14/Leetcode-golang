package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums       []int
	k          int
	multiplier int
}

type question3264 struct {
	params
	ans []int
}

func Test_Problem3264(t *testing.T) {
	qs := []question3264{
		{
			params{nums: []int{2, 1, 3, 5, 6}, k: 5, multiplier: 2},
			[]int{8, 4, 6, 5, 6},
		},
		{
			params{nums: []int{1, 2}, k: 3, multiplier: 4},
			[]int{16, 8},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, getFinalState(p.nums, p.k, p.multiplier), ans)
	}
}
