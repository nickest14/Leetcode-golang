package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2762 struct {
	params
	ans int64
}

func Test_Problem2762(t *testing.T) {
	qs := []question2762{
		{
			params{[]int{5, 4, 2, 4}},
			8,
		},
		{
			params{[]int{1, 2, 3}},
			6,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, continuousSubarrays(p.nums), ans)
	}
}
