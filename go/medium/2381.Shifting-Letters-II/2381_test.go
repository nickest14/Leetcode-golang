package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s      string
	shifts [][]int
}

type question2381 struct {
	params
	ans string
}

func Test_Problem2381(t *testing.T) {
	qs := []question2381{
		{
			params{s: "abc", shifts: [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}},
			"ace",
		},
		{
			params{s: "dztz", shifts: [][]int{{0, 0, 0}, {1, 1, 1}}},
			"catz",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, shiftingLetters(p.s, p.shifts), ans)
	}
}
