package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3354 struct {
	params
	ans int
}

func Test_Problem3354(t *testing.T) {
	qs := []question3354{
		{
			params{nums: []int{1, 0, 2, 0, 3}},
			2,
		},
		{
			params{nums: []int{2, 3, 4, 0, 4, 1, 0}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countValidSelections(p.nums), ans)
	}
}
