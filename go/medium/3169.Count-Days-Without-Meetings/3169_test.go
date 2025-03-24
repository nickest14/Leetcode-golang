package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	days     int
	meetings [][]int
}

type question3169 struct {
	params
	ans int
}

func Test_Problem3169(t *testing.T) {
	qs := []question3169{
		{
			params{days: 10, meetings: [][]int{{5, 7}, {1, 3}, {9, 10}}},
			2,
		},
		{
			params{days: 5, meetings: [][]int{{2, 4}, {1, 3}}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countDays(p.days, p.meetings), ans)
	}
}
