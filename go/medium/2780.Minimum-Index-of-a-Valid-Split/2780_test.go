package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2780 struct {
	params
	ans int
}

func Test_Problem2780(t *testing.T) {
	qs := []question2780{
		{
			params{nums: []int{1, 2, 2, 2}},
			2,
		},
		{
			params{nums: []int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}},
			4,
		},
		{
			params{nums: []int{3, 3, 3, 3, 7, 2, 2}},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumIndex(p.nums), ans)
	}
}
