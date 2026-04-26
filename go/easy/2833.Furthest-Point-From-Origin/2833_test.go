package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	moves string
}

type question2833 struct {
	params
	ans int
}

func Test_Problem2833(t *testing.T) {
	qs := []question2833{
		{
			params{moves: "L_RL__R"},
			3,
		},
		{
			params{moves: "_R__LL_"},
			5,
		},
		{
			params{moves: "_______"},
			7,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, furthestDistanceFromOrigin(p.moves), ans)
	}
}
