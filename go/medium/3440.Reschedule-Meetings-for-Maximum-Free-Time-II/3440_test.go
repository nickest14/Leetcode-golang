package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	eventTime int
	startTime []int
	endTime   []int
}

type question3440 struct {
	params
	ans int
}

func Test_Problem(t *testing.T) {
	qs := []question3440{
		{
			params{eventTime: 5, startTime: []int{1, 3}, endTime: []int{2, 5}},
			2,
		},
		{
			params{eventTime: 10, startTime: []int{0, 7, 9}, endTime: []int{1, 8, 10}},
			7,
		},
		{
			params{eventTime: 10, startTime: []int{0, 3, 7, 9}, endTime: []int{1, 4, 8, 10}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxFreeTime(p.eventTime, p.startTime, p.endTime), ans)
	}
}
