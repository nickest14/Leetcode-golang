package leetcode

import (
	"fmt"
	"testing"
)

type params1848 struct {
	nums   []int
	target int
	start  int
}

type question1848 struct {
	params1848
	ans int
}

func Test_Problem1848(t *testing.T) {
	qs := []question1848{
		{
			params1848{nums: []int{1, 2, 3, 4, 5}, target: 5, start: 3},
			1,
		},
		{
			params1848{nums: []int{1}, target: 1, start: 0},
			0,
		},
		{
			params1848{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, target: 1, start: 0},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params1848, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, getMinDistance(p.nums, p.target, p.start), ans)
	}
}
