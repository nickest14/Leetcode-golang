package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question2966 struct {
	params
	ans [][]int
}

func Test_Problem2966(t *testing.T) {
	qs := []question2966{
		{
			params{nums: []int{1, 3, 4, 8, 7, 9, 3, 5, 1}, k: 2},
			[][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}},
		},
		{
			params{nums: []int{2, 4, 2, 2, 5, 2}, k: 2},
			[][]int{},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, divideArray(p.nums, p.k), ans)
	}
}
