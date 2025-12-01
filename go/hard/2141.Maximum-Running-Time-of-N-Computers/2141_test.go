package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n         int
	batteries []int
}

type question2141 struct {
	params
	ans int
}

func Test_Problem2141(t *testing.T) {
	qs := []question2141{
		{
			params{n: 2, batteries: []int{3, 3, 3}},
			4,
		},
		{
			params{n: 2, batteries: []int{1, 1, 1, 1}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxRunTime(p.n, p.batteries), ans)
	}
}
