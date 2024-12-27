package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	values []int
}

type question1014 struct {
	params
	ans int
}

func Test_Problem1014(t *testing.T) {
	qs := []question1014{
		{
			params{[]int{8, 1, 5, 2, 6}},
			11,
		},
		{
			params{[]int{1, 2}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxScoreSightseeingPair(p.values), ans)
	}
}
