package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2501 struct {
	params
	ans int
}

func Test_Problem2501(t *testing.T) {
	qs := []question2501{
		{
			params{[]int{4, 3, 6, 16, 8, 2}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, longestSquareStreak(p.nums), ans)
	}
}
