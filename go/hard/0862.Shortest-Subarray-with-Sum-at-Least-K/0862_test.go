package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question862 struct {
	params
	ans int
}

func Test_Problem862(t *testing.T) {
	qs := []question862{
		// {
		// 	params{nums: []int{2, -1, 2}, k: 3},
		// 	3,
		// },
		{
			params{nums: []int{0, 0, 1, 1, 2, 5, -1, 6}, k: 9},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, shortestSubarray(p.nums, p.k), ans)
	}
}
