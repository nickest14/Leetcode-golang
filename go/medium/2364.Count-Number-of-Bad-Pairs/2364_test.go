package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2364 struct {
	params
	ans int
}

func Test_Problem2364(t *testing.T) {
	qs := []question2364{
		{
			params{[]int{4, 1, 3, 3}},
			5,
		},
		{
			params{[]int{1, 2, 3, 4, 5}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countBadPairs(p.nums), ans)
	}
}
