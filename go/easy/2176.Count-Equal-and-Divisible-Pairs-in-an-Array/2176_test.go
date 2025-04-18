package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question2176 struct {
	params
	ans int
}

func Test_Problem2176(t *testing.T) {
	qs := []question2176{
		{
			params{nums: []int{3, 1, 2, 2, 2, 1, 3}, k: 2},
			4,
		},
		{
			params{nums: []int{1, 2, 3, 4}, k: 1},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countPairs(p.nums, p.k), ans)
	}
}
