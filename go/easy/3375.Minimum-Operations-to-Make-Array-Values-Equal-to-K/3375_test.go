package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question3375 struct {
	params
	ans int
}

func Test_Problem3375(t *testing.T) {
	qs := []question3375{
		{
			params{nums: []int{5, 2, 5, 4, 5}, k: 2},
			2,
		},
		{
			params{nums: []int{2, 1, 2}, k: 2},
			-1,
		},
		{
			params{nums: []int{9, 7, 5, 3}, k: 1},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minOperations(p.nums, p.k), ans)
	}
}
