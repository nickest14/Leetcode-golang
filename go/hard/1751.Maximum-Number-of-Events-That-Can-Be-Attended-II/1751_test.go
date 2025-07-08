package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	events [][]int
	k      int
}

type question1751 struct {
	params
	ans int
}

func Test_Problem1751(t *testing.T) {
	qs := []question1751{
		{
			params{events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}, k: 2},
			7,
		},
		{
			params{events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}}, k: 2},
			10,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxValue(p.events, p.k), ans)
	}
}
