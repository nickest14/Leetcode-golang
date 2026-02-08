package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question3634 struct {
	params
	ans int
}

func Test_Problem(t *testing.T) {
	qs := []question3634{
		{
			params{nums: []int{2, 1, 5}, k: 2},
			1,
		},
		{
			params{nums: []int{1, 6, 2, 9}, k: 3},
			2,
		},
		{
			params{nums: []int{4, 6}, k: 2},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minRemoval(p.nums, p.k), ans)
	}
}
