package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	events [][]int
}

type question1353 struct {
	params
	ans int
}

func Test_Problem1353(t *testing.T) {
	qs := []question1353{
		{
			params{[][]int{{1, 2}, {2, 3}, {3, 4}}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxEvents(p.events), ans)
	}
}
