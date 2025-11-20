package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums     []int
	original int
}

type question2154 struct {
	params
	ans int
}

func Test_Problem2154(t *testing.T) {
	qs := []question2154{
		{
			params{nums: []int{5, 3, 6, 1, 12}, original: 3},
			24,
		},
		{
			params{nums: []int{2, 7, 9}, original: 4},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findFinalValue(p.nums, p.original), ans)
	}
}
