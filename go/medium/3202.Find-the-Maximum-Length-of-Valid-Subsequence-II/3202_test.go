package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}
type question3202 struct {
	params
	ans int
}

func Test_Problem3202(t *testing.T) {
	qs := []question3202{
		{
			params{nums: []int{1, 2, 3, 4, 5}, k: 2},
			5,
		},
		{
			params{nums: []int{1, 4, 2, 3, 1, 4}, k: 3},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumLength(p.nums, p.k), ans)
	}
}
