package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums   []int
	target int
}

type question1498 struct {
	params
	ans int
}

func Test_Problem498(t *testing.T) {
	qs := []question1498{
		{
			params{nums: []int{3, 5, 6, 7}, target: 9},
			4,
		},
		{
			params{nums: []int{3, 3, 6, 8}, target: 10},
			6,
		},
		{
			params{nums: []int{2, 3, 3, 4, 6, 7}, target: 12},
			61,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numSubseq(p.nums, p.target), ans)
	}
}
