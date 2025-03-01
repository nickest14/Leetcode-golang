package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question962 struct {
	params
	ans int
}

func Test_Problem962(t *testing.T) {
	qs := []question962{
		{
			params{nums: []int{6, 0, 8, 2, 1, 5}},
			4,
		},
		{
			params{nums: []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}},
			7,
		},
		{
			params{nums: []int{3, 28, 15, 1, 4, 12, 6, 19, 8, 15, 3, 9, 6, 4, 13, 12, 6, 12, 10, 1, 2, 1, 4, 1, 4, 0, 0, 1, 1, 0}},
			25,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxWidthRamp(p.nums), ans)
	}
}
