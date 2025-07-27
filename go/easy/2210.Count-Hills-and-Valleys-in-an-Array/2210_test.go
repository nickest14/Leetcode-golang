package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2210 struct {
	params
	ans int
}

func Test_Problem2210(t *testing.T) {
	qs := []question2210{
		{
			params{[]int{2, 4, 1, 1, 6, 5}},
			3,
		},
		{
			params{[]int{6, 6, 5, 5, 4, 1}},
			0,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countHillValley(p.nums), ans)
	}
}
