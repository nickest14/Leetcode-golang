package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question912 struct {
	params
	ans []int
}

func Test_Problem912(t *testing.T) {
	qs := []question912{
		{
			params{nums: []int{5, 1, 1, 2, 0, 0}},
			[]int{0, 0, 1, 1, 2, 5},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, sortArray(p.nums), ans)
	}
}
