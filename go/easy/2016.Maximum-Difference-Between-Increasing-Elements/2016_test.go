package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2016 struct {
	params
	ans int
}

func Test_Problem2016(t *testing.T) {
	qs := []question2016{
		{
			params{[]int{7, 1, 5, 4}},
			4,
		},
		{
			params{[]int{9, 4, 3, 2}},
			-1,
		},
		{
			params{[]int{1, 5, 2, 10}},
			9,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumDifference(p.nums), ans)
	}
}
