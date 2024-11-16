package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question3254 struct {
	params
	ans []int
}

func Test_Problem3254(t *testing.T) {
	qs := []question3254{
		{
			params{nums: []int{1, 2, 3, 4, 3, 2, 5}, k: 3},
			[]int{3, 4, -1, -1, -1},
		},
		{
			params{nums: []int{3, 2, 3, 2, 3, 2}, k: 2},
			[]int{-1, 3, -1, 3, -1},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, resultsArray(p.nums, p.k), ans)
	}
}
