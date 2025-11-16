package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3289 struct {
	params
	ans []int
}

func Test_Problem3289(t *testing.T) {
	qs := []question3289{
		{
			params{nums: []int{0, 1, 1, 0}},
			[]int{0, 1},
		},
		{
			params{nums: []int{0, 3, 2, 1, 3, 2}},
			[]int{2, 3},
		},
		{
			params{nums: []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}},
			[]int{4, 5},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, getSneakyNumbers(p.nums), ans)
	}
}
