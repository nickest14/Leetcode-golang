package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums  []int
	value int
}

type question2598 struct {
	params
	ans int
}

func Test_Problem2598(t *testing.T) {
	qs := []question2598{
		{
			params{nums: []int{1, -10, 7, 13, 6, 8}, value: 4},
			4,
		},
		{
			params{nums: []int{1, -10, 7, 13, 6, 8}, value: 7},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findSmallestInteger(p.nums, p.value), ans)
	}
}
