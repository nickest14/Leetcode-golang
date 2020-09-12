package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums   []int
	target int
}

type ans struct {
	li []int
}

type question struct {
	params
	ans
}

// TestProblem1 is to test problem 1
func TestProblem1(t *testing.T) {

	qs := []question{
		{
			params{[]int{1, 2, 3}, 5},
			ans{[]int{1, 2}},
		},
		{
			params{[]int{4, 2, 9, 1}, 3},
			ans{[]int{1, 3}},
		},
		{
			params{[]int{7, 1, 5, 2, 9}, 16},
			ans{[]int{0, 4}},
		},
	}

	for _, q := range qs {
		p, _ := q.params, q.ans
		ans := TwoSum(p.nums, p.target)
		fmt.Printf("[input]: %v        [output]: %v\n", p, ans)
	}
}
