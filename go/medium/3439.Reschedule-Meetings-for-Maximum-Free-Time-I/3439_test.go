package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	eventTime int
	k         int
	startTime []int
	endTime   []int
}

type question3439 struct {
	params
	ans int
}

func Test_Problem(t *testing.T) {
	qs := []question3439{
		{
			params{eventTime: 5, k: 1, startTime: []int{1, 3}, endTime: []int{2, 5}},
			2,
		},
		{
			params{eventTime: 10, k: 1, startTime: []int{0, 2, 9}, endTime: []int{1, 4, 10}},
			6,
		},
		{
			params{eventTime: 5, k: 2, startTime: []int{0, 1, 2, 3, 4}, endTime: []int{1, 2, 3, 4, 5}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxFreeTime(p.eventTime, p.k, p.startTime, p.endTime), ans)
	}
}
