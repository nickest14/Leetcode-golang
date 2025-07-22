package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1695 struct {
	params
	ans int
}

func Test_Problem(t *testing.T) {
	qs := []question1695{
		{
			params{nums: []int{4, 2, 4, 5, 6}},
			17,
		},
		{
			params{nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5}},
			8,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumUniqueSubarray(p.nums), ans)
	}
}
